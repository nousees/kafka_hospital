package kafka

import (
	"api_service/internal/entities"
	"encoding/json"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(broker string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		logrus.Errorf("Failed to create Sarama producer: %v", err)
		return nil, err
	}

	logrus.Info("Sarama producer created successfully")
	return &KafkaProducer{producer: producer}, nil
}

func (p *KafkaProducer) SendAppointment(appt entities.Appointment) error {
	apptBytes, err := json.Marshal(appt)
	if err != nil {
		logrus.Errorf("Failed to marshal appointment: %v", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "appointments",
		Key:   sarama.StringEncoder(strconv.Itoa(appt.PatientID)),
		Value: sarama.StringEncoder(apptBytes),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		logrus.Errorf("Failed to send appointment to Kafka: %v", err)
		return err
	}

	logrus.Infof("Appointment sent to Kafka: partition=%d, offset=%d", partition, offset)
	return nil
}

func (p *KafkaProducer) SendPatient(patient entities.Patient) error {
	patientBytes, err := json.Marshal(patient)
	if err != nil {
		logrus.Errorf("Failed to marshal patient: %v", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "patients",
		Key:   sarama.StringEncoder(patient.Name),
		Value: sarama.StringEncoder(patientBytes),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		logrus.Errorf("Failed to send patient to Kafka: %v", err)
		return err
	}

	logrus.Infof("Patient sent to Kafka: partition=%d, offset=%d", partition, offset)
	return nil
}

func (p *KafkaProducer) SendDoctor(doctor entities.Doctor) error {
	doctorBytes, err := json.Marshal(doctor)
	if err != nil {
		logrus.Errorf("Failed to marshal doctor: %v", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "doctors",
		Key:   sarama.StringEncoder(doctor.Name),
		Value: sarama.StringEncoder(doctorBytes),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		logrus.Errorf("Failed to send doctor to Kafka: %v", err)
		return err
	}

	logrus.Infof("Doctor sent to Kafka: partition=%d, offset=%d", partition, offset)
	return nil
}

func (p *KafkaProducer) Close() {
	if err := p.producer.Close(); err != nil {
		logrus.Errorf("Failed to close Sarama producer: %v", err)
	}
	logrus.Info("Sarama producer closed")
}
