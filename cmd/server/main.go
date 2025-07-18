package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/talanayush/gofr-chat/docs" // Swagger docs import
	"github.com/talanayush/gofr-chat/internal/handler"
)

func main() {
	r := mux.NewRouter()

	// Swagger UI route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Health check
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	log.Println("Server running at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
