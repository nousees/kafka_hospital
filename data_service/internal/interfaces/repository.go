package interfaces

import "data_service/internal/entities"

type AppointmentRepository interface {
	Save(appt entities.Appointment) error
	SavePatient(patient entities.Patient) error
	SaveDoctor(doctor entities.Doctor) error
	Search() ([]map[string]interface{}, error)
	GetDailyCounts() ([]map[string]interface{}, error)
	GetTopPatients() ([]map[string]interface{}, error)
	GetTopDoctors() ([]map[string]interface{}, error)
}
