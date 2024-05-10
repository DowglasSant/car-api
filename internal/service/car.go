package service

import (
	"car-api/internal/repository"
	"car-api/pkg/model"
	"context"
	"log"
)

type CarService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) GetCars(ctx context.Context) ([]model.Car, error) {
	cars, err := s.repo.GetCars(ctx)
	if err != nil {
		log.Printf("Error retrieving cars: %v", err)
		return nil, err
	}
	return cars, nil
}

func (s *CarService) GetCarByID(ctx context.Context, id uint) (model.Car, error) {
	car, err := s.repo.GetCarByID(ctx, id)
	if err != nil {
		log.Printf("Error retrieving car with ID %d: %v", id, err)
		return model.Car{}, err
	}
	return car, nil
}

func (s *CarService) CreateCar(ctx context.Context, car model.Car) error {
	err := s.repo.CreateCar(ctx, car)
	if err != nil {
		log.Printf("Error creating new car: %v", err)
		return err
	}
	return nil
}

func (s *CarService) UpdateCar(ctx context.Context, car model.Car) error {
	err := s.repo.UpdateCar(ctx, car)
	if err != nil {
		log.Printf("Error updating car with ID %d: %v", car.ID, err)
		return err
	}
	return nil
}

func (s *CarService) DeleteCar(ctx context.Context, id uint) error {
	err := s.repo.DeleteCar(ctx, id)
	if err != nil {
		log.Printf("Error deleting car with ID %d: %v", id, err)
		return err
	}
	return nil
}

func (s *CarService) GetCarByLicensePlate(ctx context.Context, licensePlate string) (model.Car, error) {
	car, err := s.repo.GetCarByLicensePlate(ctx, licensePlate)
	if err != nil {
		log.Printf("Error retrieving car with license plate %s: %v", licensePlate, err)
	}
	return car, err
}

func (s *CarService) GetCarsByMake(ctx context.Context, make string) ([]model.Car, error) {
	cars, err := s.repo.GetCarsByMake(ctx, make)
	if err != nil {
		log.Printf("Error retrieving cars by make %s: %v", make, err)
	}
	return cars, err
}
