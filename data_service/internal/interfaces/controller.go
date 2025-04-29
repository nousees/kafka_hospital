package interfaces

import (
	"net/http"

	inftHTTP "data_service/internal/infrastructure/http"
)

type HTTPController struct {
	usecase AppointmentUsecase
}

type AppointmentUsecase interface {
	SearchAppointments() ([]map[string]interface{}, error)
	GetReport(reportType string) ([]map[string]interface{}, error)
}

func NewHTTPController(usecase AppointmentUsecase) *HTTPController {
	return &HTTPController{usecase: usecase}
}

func (c *HTTPController) SearchAppointments(w http.ResponseWriter, r *http.Request) {
	results, err := c.usecase.SearchAppointments()
	if err != nil {
		http.Error(w, "Failed to fetch appointments: "+err.Error(), http.StatusInternalServerError)
		return
	}
	inftHTTP.JSONResponse(w, results)
}

func (c *HTTPController) GetReport(w http.ResponseWriter, r *http.Request) {
	reportType := inftHTTP.GetReportType(r)
	results, err := c.usecase.GetReport(reportType)
	if err != nil {
		http.Error(w, "Failed to fetch report: "+err.Error(), http.StatusBadRequest)
		return
	}
	inftHTTP.JSONResponse(w, results)
}
