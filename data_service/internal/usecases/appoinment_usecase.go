package usecases

import (
	"data_service/internal/entities"
	"data_service/internal/interfaces"
	"fmt"
)

type AppointmentUsecase struct {
	repo interfaces.AppointmentRepository
}

func NewAppointmentUsecase(repo interfaces.AppointmentRepository) *AppointmentUsecase {
	return &AppointmentUsecase{repo: repo}
}

func (u *AppointmentUsecase) SaveAppointment(appt entities.Appointment) error {
	return u.repo.Save(appt)
}

func (u *AppointmentUsecase) SavePatient(patient entities.Patient) error {
	return u.repo.SavePatient(patient)
}

func (u *AppointmentUsecase) SaveDoctor(doctor entities.Doctor) error {
	return u.repo.SaveDoctor(doctor)
}

func (u *AppointmentUsecase) SearchAppointments() ([]map[string]interface{}, error) {
	return u.repo.Search()
}

func (u *AppointmentUsecase) GetReport(reportType string) ([]map[string]interface{}, error) {
	switch reportType {
	case "daily-counts":
		return u.repo.GetDailyCounts()
	case "top-patients":
		return u.repo.GetTopPatients()
	case "top-doctors":
		return u.repo.GetTopDoctors()
	default:
		return nil, fmt.Errorf("invalid report type")
	}
}
