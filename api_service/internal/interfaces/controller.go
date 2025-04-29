package interfaces

import (
	"encoding/json"
	"net/http"

	"api_service/internal/entities"
	infrHTTP "api_service/internal/infrastructure/http"
)

type HTTPController struct {
	usecase AppointmentUsecase
}

type AppointmentUsecase interface {
	CreateAppointment(appt entities.Appointment) error
	CreatePatient(patient entities.Patient) error
	CreateDoctor(doctor entities.Doctor) error
	SearchAppointments() ([]map[string]interface{}, error)
	GetReport(reportType string) ([]map[string]interface{}, error)
}

func NewHTTPController(usecase AppointmentUsecase) *HTTPController {
	return &HTTPController{usecase: usecase}
}

func (c *HTTPController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appt entities.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.usecase.CreateAppointment(appt); err != nil {
		http.Error(w, "Failed to create appointment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Appointment created"})
}

func (c *HTTPController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient entities.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.usecase.CreatePatient(patient); err != nil {
		http.Error(w, "Failed to create patient: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Patient created"})
}

func (c *HTTPController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor entities.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.usecase.CreateDoctor(doctor); err != nil {
		http.Error(w, "Failed to create doctor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Doctor created"})
}

func (c *HTTPController) SearchAppointments(w http.ResponseWriter, r *http.Request) {
	results, err := c.usecase.SearchAppointments()
	if err != nil {
		http.Error(w, "Failed to fetch appointments: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

func (c *HTTPController) GetReport(w http.ResponseWriter, r *http.Request) {
	reportType := infrHTTP.GetReportType(r)
	results, err := c.usecase.GetReport(reportType)
	if err != nil {
		http.Error(w, "Failed to fetch report: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}
