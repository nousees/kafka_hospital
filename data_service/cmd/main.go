package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"data_service/internal/infrastructure/db"
	infrHTTP "data_service/internal/infrastructure/http"
	"data_service/internal/infrastructure/kafka"
	"data_service/internal/interfaces"
	"data_service/internal/usecases"

	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	kafkaConsumer, err := kafka.NewKafkaConsumer(os.Getenv("KAFKA_BROKER"))
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	dbRepo := db.NewPostgresRepository(dbConn)
	appointmentUsecase := usecases.NewAppointmentUsecase(dbRepo)
	controller := interfaces.NewHTTPController(appointmentUsecase)

	go kafka.ConsumeAppointments(kafkaConsumer, dbRepo)

	router := infrHTTP.NewRouter()
	router.HandleSearchAppointments(controller)
	router.HandleReports(controller)

	log.Println("Data Service running on :8081")
	http.ListenAndServe(":8081", router)
}
