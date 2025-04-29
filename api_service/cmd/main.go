package main

import (
	"log"
	"net/http"
	"os"

	infrHTTP "api_service/internal/infrastructure/http"
	"api_service/internal/infrastructure/kafka"
	"api_service/internal/interfaces"
	"api_service/internal/usecases"
)

func main() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaProducer, err := kafka.NewKafkaProducer(kafkaBroker)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	appointmentUsecase := usecases.NewAppointmentUsecase(kafkaProducer)
	controller := interfaces.NewHTTPController(appointmentUsecase)

	router := infrHTTP.NewRouter()
	router.HandleCreateAppointment(controller)
	router.HandleCreatePatient(controller)
	router.HandleCreateDoctor(controller)
	router.HandleSearchAppointments(controller)
	router.HandleReports(controller)

	log.Println("API Service running on :8080")
	http.ListenAndServe(":8080", router)
}
