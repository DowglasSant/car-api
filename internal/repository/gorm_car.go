package repository

import (
	"car-api/pkg/model"
	"context"

	"gorm.io/gorm"
)

type gormCarRepository struct {
	db *gorm.DB
}

func NewGormCarRepository(db *gorm.DB) CarRepository {
	return &gormCarRepository{db: db}
}

func (r *gormCarRepository) GetCars(ctx context.Context) ([]model.Car, error) {
	var cars []model.Car
	result := r.db.Find(&cars)
	return cars, result.Error
}

func (r *gormCarRepository) GetCarByID(ctx context.Context, id uint) (model.Car, error) {
	var car model.Car
	result := r.db.First(&car, id)
	return car, result.Error
}

func (r *gormCarRepository) GetCarByLicensePlate(ctx context.Context, licensePlate string) (model.Car, error) {
	var car model.Car
	result := r.db.Where("license_plate = ?", licensePlate).First(&car)
	return car, result.Error
}

func (r *gormCarRepository) GetCarsByMake(ctx context.Context, make string) ([]model.Car, error) {
	var cars []model.Car
	result := r.db.Where("make = ?", make).Find(&cars)
	return cars, result.Error
}

func (r *gormCarRepository) CreateCar(ctx context.Context, car model.Car) error {
	result := r.db.Create(&car)
	return result.Error
}

func (r *gormCarRepository) UpdateCar(ctx context.Context, car model.Car) error {
	result := r.db.Save(&car)
	return result.Error
}

func (r *gormCarRepository) DeleteCar(ctx context.Context, id uint) error {
	result := r.db.Delete(&model.Car{}, id)
	return result.Error
}
