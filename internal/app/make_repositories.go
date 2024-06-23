package app

import (
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/repository"
	"time"
)

type connectionCloser interface {
	Close()
}

type orderRepository interface {
	connectionCloser
	Push(order model.Order) []model.Order
}

type availabilityRepository interface {
	connectionCloser
	GetAll() []model.RoomAvailability
	Replace(index int, record model.RoomAvailability) error
	DecrementQuotaForRange(from, to time.Time)
}

type repositories struct {
	OrderRepository        orderRepository
	AvailabilityRepository availabilityRepository
}

func makeRepositories() repositories {
	availabilityRepository := repository.NewAvailabilityRepository()
	availabilityRepository.LoadFixtures()
	return repositories{
		OrderRepository:        repository.NewOrderRepository(),
		AvailabilityRepository: availabilityRepository,
	}
}
