package kafka

import (
	"data_service/internal/entities"
	"data_service/internal/interfaces"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func NewKafkaConsumer(broker string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		logrus.Errorf("Failed to create Sarama consumer: %v", err)
		return nil, err
	}

	logrus.Info("Sarama consumer created successfully")
	return consumer, nil
}

func entityHandler[T any](save func(T) error) func([]byte) error {
	return func(value []byte) error {
		var entity T
		if err := json.Unmarshal(value, &entity); err != nil {
			return err
		}
		return save(entity)
	}
}

func createHandlers(repo interfaces.AppointmentRepository) map[string]func([]byte) error {
	return map[string]func([]byte) error{
		"appointments": entityHandler[entities.Appointment](repo.Save),
		"patients":     entityHandler[entities.Patient](repo.SavePatient),
		"doctors":      entityHandler[entities.Doctor](repo.SaveDoctor),
	}
}

func ConsumeAppointments(consumer sarama.Consumer, repo interfaces.AppointmentRepository) {
	handlers := createHandlers(repo)

	for topic, handler := range handlers {
		partitions, err := consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Failed to get partitions for topic %s: %v", topic, err)
			continue
		}

		for _, partition := range partitions {
			pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
			if err != nil {
				logrus.Errorf("Failed to consume partition %d for topic %s: %v", partition, topic, err)
				continue
			}

			go func(topic string, pc sarama.PartitionConsumer, handler func([]byte) error) {
				for msg := range pc.Messages() {
					if err := handler(msg.Value); err != nil {
						logrus.Errorf("Error handling message from topic %s: %v", topic, err)
					}
				}
			}(topic, pc, handler)
		}
	}

	select {}
}
