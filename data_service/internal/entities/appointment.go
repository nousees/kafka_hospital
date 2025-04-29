package entities

type Appointment struct {
	PatientID       int    `json:"patient_id"`
	DoctorID        int    `json:"doctor_id"`
	AppointmentDate string `json:"appointment_date"`
}
