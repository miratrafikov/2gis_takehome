package createorder

import (
	"applicationDesignTest/internal/log"
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/repository"
	"errors"
	"time"
)

type Usecase struct {
	orderRepository        repository.OrderRepository
	availabilityRepository repository.AvailabilityRepository
	logger                 log.Logger
}

func New(
	orderRepository repository.OrderRepository,
	availabilityRepository repository.AvailabilityRepository,
	logger log.Logger,
) Usecase {
	return Usecase{
		orderRepository: orderRepository,
		logger:          logger,
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

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
