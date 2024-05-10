package main

import (
	"car-api/internal/app"
	"car-api/internal/repository"
	"car-api/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	carRepo := repository.NewGormCarRepository(db)
	carService := service.NewCarService(carRepo)

	r := mux.NewRouter()
	app.RegisterCarRoutes(r, carService)

	log.Println("Starting server on :8085")
	http.ListenAndServe(":8085", r)
}
