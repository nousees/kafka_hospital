package db

import (
	"data_service/internal/entities"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(appt entities.Appointment) error {
	if appt.AppointmentDate == "" {
		logrus.Errorf("Appointment date is empty")
	}
	_, err := r.db.Exec("INSERT INTO appointments (patient_id, doctor_id, appointment_date) VALUES ($1, $2, $3)", appt.PatientID, appt.DoctorID, appt.AppointmentDate)
	if err != nil {
		logrus.Errorf("Failed to save appointment: %v", err)
		return err
	}
	logrus.Info("Appointment saved successfully")
	return nil
}

func (r *PostgresRepository) SavePatient(patient entities.Patient) error {
	_, err := r.db.Exec("INSERT INTO patients (name) VALUES ($1)", patient.Name)
	if err != nil {
		logrus.Errorf("Failed to save patient: %v", err)
		return err
	}
	logrus.Info("Patient saved successfully")
	return nil
}

func (r *PostgresRepository) SaveDoctor(doctor entities.Doctor) error {
	_, err := r.db.Exec("INSERT INTO doctors (name) VALUES ($1)", doctor.Name)
	if err != nil {
		logrus.Errorf("Failed to save doctor: %v", err)
		return err
	}
	logrus.Info("Doctor saved successfully")
	return nil
}

func (r *PostgresRepository) Search() ([]map[string]interface{}, error) {
	rows, err := r.db.Query("SELECT a.id, a.patient_id, a.doctor_id, a.appointment_date, p.name as patient_name, d.name as doctor_name FROM appointments a JOIN patients p ON a.patient_id = p.id JOIN doctors d ON a.doctor_id = d.id")
	if err != nil {
		logrus.Errorf("Failed to search appointments: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id, patientID, doctorID int
		var appointmentDate, patientName, doctorName string
		if err := rows.Scan(&id, &patientID, &doctorID, &appointmentDate, &patientName, &doctorName); err != nil {
			logrus.Errorf("Failed to scan appointment row: %v", err)
			return nil, err
		}
		results = append(results, map[string]interface{}{
			"id":               id,
			"patient_name":     patientName,
			"doctor_name":      doctorName,
			"appointment_date": appointmentDate,
		})
	}
	logrus.Info("Appointments retrieved successfully")
	return results, nil
}

func (r *PostgresRepository) GetDailyCounts() ([]map[string]interface{}, error) {
	rows, err := r.db.Query("SELECT appointment_date, COUNT(*) as count FROM appointments GROUP BY appointment_date")
	if err != nil {
		logrus.Errorf("Failed to get daily counts: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var date string
		var count int
		if err := rows.Scan(&date, &count); err != nil {
			logrus.Errorf("Failed to scan daily counts row: %v", err)
			return nil, err
		}
		results = append(results, map[string]interface{}{"date": date, "count": count})
	}
	logrus.Info("Daily counts retrieved successfully")
	return results, nil
}

func (r *PostgresRepository) GetTopPatients() ([]map[string]interface{}, error) {
	rows, err := r.db.Query("SELECT p.name, COUNT(*) as count FROM appointments a JOIN patients p ON a.patient_id = p.id GROUP BY p.id, p.name ORDER BY count DESC LIMIT 10")
	if err != nil {
		logrus.Errorf("Failed to get top patients: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var name string
		var count int
		if err := rows.Scan(&name, &count); err != nil {
			logrus.Errorf("Failed to scan top patients row: %v", err)
			return nil, err
		}
		results = append(results, map[string]interface{}{"name": name, "count": count})
	}
	logrus.Info("Top patients retrieved successfully")
	return results, nil
}

func (r *PostgresRepository) GetTopDoctors() ([]map[string]interface{}, error) {
	rows, err := r.db.Query("SELECT d.name, COUNT(*) as count FROM appointments a JOIN doctors d ON a.doctor_id = d.id GROUP BY d.id, d.name ORDER BY count DESC LIMIT 10")
	if err != nil {
		logrus.Errorf("Failed to get top doctors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var name string
		var count int
		if err := rows.Scan(&name, &count); err != nil {
			logrus.Errorf("Failed to scan top doctors row: %v", err)
			return nil, err
		}
		results = append(results, map[string]interface{}{"name": name, "count": count})
	}
	logrus.Info("Top doctors retrieved successfully")
	return results, nil
}
