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

func ConsumeAppointments(consumer sarama.Consumer, repo interfaces.AppointmentRepository) {
	topics := []string{"appointments", "patients", "doctors"}

	for _, topic := range topics {
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

			go func(topic string, pc sarama.PartitionConsumer) {
				for msg := range pc.Messages() {
					switch topic {
					case "appointments":
						var appt entities.Appointment
						if err := json.Unmarshal(msg.Value, &appt); err != nil {
							logrus.Errorf("Failed to unmarshal appointment: %v", err)
							continue
						}
						if err := repo.Save(appt); err != nil {
							logrus.Errorf("Failed to save appointment: %v", err)
						}
					case "patients":
						var patient entities.Patient
						if err := json.Unmarshal(msg.Value, &patient); err != nil {
							logrus.Errorf("Failed to unmarshal patient: %v", err)
							continue
						}
						if err := repo.SavePatient(patient); err != nil {
							logrus.Errorf("Failed to save patient: %v", err)
						}
					case "doctors":
						var doctor entities.Doctor
						if err := json.Unmarshal(msg.Value, &doctor); err != nil {
							logrus.Errorf("Failed to unmarshal doctor: %v", err)
							continue
						}
						if err := repo.SaveDoctor(doctor); err != nil {
							logrus.Errorf("Failed to save doctor: %v", err)
						}
					}
				}
			}(topic, pc)
		}
	}

	select {}
}
