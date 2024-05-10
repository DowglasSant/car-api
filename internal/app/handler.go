package app

import (
	"car-api/internal/service"
	"car-api/pkg/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type CarHandler struct {
	service *service.CarService
}

func NewCarHandler(service *service.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.service.GetCars(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cars)
}

func (h *CarHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	car, err := h.service.GetCarByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(car)
}

func (h *CarHandler) AddCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCar(r.Context(), car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCar(r.Context(), car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCar(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Car deleted successfully"))
}

func (h *CarHandler) GetCarByLicensePlate(w http.ResponseWriter, r *http.Request) {
	licensePlate := r.URL.Query().Get("license_plate")
	if licensePlate == "" {
		http.Error(w, "License plate is required", http.StatusBadRequest)
		return
	}

	car, err := h.service.GetCarByLicensePlate(r.Context(), licensePlate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(car)
}

func (h *CarHandler) GetCarsByMake(w http.ResponseWriter, r *http.Request) {
	make := r.URL.Query().Get("make")
	if make == "" {
		http.Error(w, "Make is required", http.StatusBadRequest)
		return
	}

	cars, err := h.service.GetCarsByMake(r.Context(), make)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cars)
}
