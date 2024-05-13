package repository

import "applicationDesignTest/internal/model"

type Repository interface {
	Close()
}

type OrderRepository interface {
	Repository
	Push(order model.Order) []model.Order
}

type AvailabilityRepository interface {
	Repository
	Push(order model.RoomAvailability) []model.RoomAvailability
	GetByIndex(index int) (model.RoomAvailability, error)
	GetAll() []model.RoomAvailability
	Replace(index int, record model.RoomAvailability) error
}
