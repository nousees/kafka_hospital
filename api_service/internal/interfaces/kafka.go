package interfaces

import "api_service/internal/entities"

type KafkaProducer interface {
	SendAppointment(appt entities.Appointment) error
	SendPatient(patient entities.Patient) error
	SendDoctor(doctor entities.Doctor) error
	Close()
}
