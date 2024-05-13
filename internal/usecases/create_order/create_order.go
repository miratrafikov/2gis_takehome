package createorder

import (
	"applicationDesignTest/internal/log"
	"applicationDesignTest/internal/model"
	"errors"
	"time"
)

type OrderRepository interface {
	Push(order model.Order) []model.Order
}

type AvailabilityRepository interface {
	GetAll() []model.RoomAvailability
	Replace(index int, record model.RoomAvailability) error
}

type Usecase struct {
	orderRepository        OrderRepository
	availabilityRepository AvailabilityRepository
	logger                 log.Logger
}

func New(
	orderRepository OrderRepository,
	availabilityRepository AvailabilityRepository,
	logger log.Logger,
) Usecase {
	return Usecase{
		orderRepository:        orderRepository,
		availabilityRepository: availabilityRepository,
		logger:                 logger,
	}
}

const roomUnavailableError = "hotel room is not available for selected dates"

func (uc Usecase) Handle(order model.Order) (model.Order, error) {
	daysToBook := getTimestampsOfDaysBetweenDates(order.From, order.To)
	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}
	for _, dayToBook := range daysToBook {
		availableDates := uc.availabilityRepository.GetAll()
		for i, availability := range availableDates {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			uc.availabilityRepository.Replace(i, availability)
			delete(unavailableDays, dayToBook)
		}
	}
	if len(unavailableDays) != 0 {
		uc.logger.Errorf("%s:\n%v\n%v", roomUnavailableError, order, unavailableDays)
		return model.Order{}, errors.New(roomUnavailableError)
	}
	uc.orderRepository.Push(order)
	uc.logger.Infof("Order successfully created: %v", order)
	return order, nil
}

func getTimestampsOfDaysBetweenDates(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}
	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
