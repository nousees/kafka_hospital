package usecases

import (
	"api_service/internal/entities"
	"api_service/internal/interfaces"
)

type AppointmentUsecase struct {
	kafkaProducer interfaces.KafkaProducer
}

func NewAppointmentUsecase(kafkaProducer interfaces.KafkaProducer) *AppointmentUsecase {
	return &AppointmentUsecase{kafkaProducer: kafkaProducer}
}

func (u *AppointmentUsecase) CreateAppointment(appt entities.Appointment) error {
	return u.kafkaProducer.SendAppointment(appt)
}

func (u *AppointmentUsecase) CreatePatient(patient entities.Patient) error {
	return u.kafkaProducer.SendPatient(patient)
}

func (u *AppointmentUsecase) CreateDoctor(doctor entities.Doctor) error {
	return u.kafkaProducer.SendDoctor(doctor)
}

func (u *AppointmentUsecase) SearchAppointments() ([]map[string]interface{}, error) {
	return interfaces.FetchFromDataService("http://data-service:8081/search")
}

func (u *AppointmentUsecase) GetReport(reportType string) ([]map[string]interface{}, error) {
	return interfaces.FetchFromDataService("http://data-service:8081/reports/" + reportType)
}
