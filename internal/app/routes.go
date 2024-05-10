package app

import (
	"car-api/internal/service"

	"github.com/gorilla/mux"
)

func RegisterCarRoutes(router *mux.Router, carService *service.CarService) {
	carHandler := NewCarHandler(carService)
	router.HandleFunc("/cars", carHandler.GetCars).Methods("GET")
	router.HandleFunc("/cars", carHandler.AddCar).Methods("POST")
	router.HandleFunc("/cars/{id:[0-9]+}", carHandler.GetCarByID).Methods("GET")
	router.HandleFunc("/cars", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id:[0-9]+}", carHandler.DeleteCar).Methods("DELETE")
	router.HandleFunc("/cars/license_plate", carHandler.GetCarByLicensePlate).Methods("GET")
	router.HandleFunc("/cars/make", carHandler.GetCarsByMake).Methods("GET")
}
