package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{router: mux.NewRouter()}
}

func (r *Router) HandleCreateAppointment(controller interface {
	CreateAppointment(w http.ResponseWriter, r *http.Request)
}) {
	r.router.HandleFunc("/appointment", controller.CreateAppointment).Methods("POST")
}

func (r *Router) HandleCreatePatient(controller interface {
	CreatePatient(w http.ResponseWriter, r *http.Request)
}) {
	r.router.HandleFunc("/patients", controller.CreatePatient).Methods("POST")
}

func (r *Router) HandleCreateDoctor(controller interface {
	CreateDoctor(w http.ResponseWriter, r *http.Request)
}) {
	r.router.HandleFunc("/doctors", controller.CreateDoctor).Methods("POST")
}

func (r *Router) HandleSearchAppointments(controller interface {
	SearchAppointments(w http.ResponseWriter, r *http.Request)
}) {
	r.router.HandleFunc("/search", controller.SearchAppointments).Methods("GET")
}

func (r *Router) HandleReports(controller interface {
	GetReport(w http.ResponseWriter, r *http.Request)
}) {
	r.router.HandleFunc("/reports/{type}", controller.GetReport).Methods("GET")
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func GetReportType(r *http.Request) string {
	return mux.Vars(r)["type"]
}
