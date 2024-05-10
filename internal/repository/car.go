package repository

import (
	"car-api/pkg/model"
	"context"
)

type CarRepository interface {
	GetCars(ctx context.Context) ([]model.Car, error)
	GetCarByID(ctx context.Context, id uint) (model.Car, error)
	GetCarByLicensePlate(ctx context.Context, licensePlate string) (model.Car, error)
	GetCarsByMake(ctx context.Context, make string) ([]model.Car, error)
	CreateCar(ctx context.Context, car model.Car) error
	UpdateCar(ctx context.Context, car model.Car) error
	DeleteCar(ctx context.Context, id uint) error
}
