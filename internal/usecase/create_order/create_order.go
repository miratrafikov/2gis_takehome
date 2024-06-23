package createorder

import (
	"applicationDesignTest/internal/log"
	"applicationDesignTest/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

type orderRepository interface {
	Push(order model.Order) []model.Order
}

type availabilityRepository interface {
	GetAll() []model.RoomAvailability
	Replace(index int, record model.RoomAvailability) error
	DecrementQuotaForRange(from, to time.Time)
}

type Usecase struct {
	orderRepository        orderRepository
	availabilityRepository availabilityRepository
	logger                 log.Logger
	bookingMutex           *sync.Mutex
}

func New(
	orderRepository orderRepository,
	availabilityRepository availabilityRepository,
	logger log.Logger,
) Usecase {
	return Usecase{
		orderRepository:        orderRepository,
		availabilityRepository: availabilityRepository,
		logger:                 logger,
	}
}

func (u Usecase) Handle(order model.Order) (model.Order, error) {
	order.From = floorDate(order.From)
	order.To = floorDate(order.To)
	if err := u.validate(order); err != nil {
		return model.Order{}, err
	}
	datesToBook := getDatesInRange(order.From, order.To)
	availabilityRecords := u.availabilityRepository.GetAll()
	dateToAvailabilityRecords := make(map[time.Time]model.RoomAvailability, len(availabilityRecords))
	for _, record := range availabilityRecords {
		dateToAvailabilityRecords[record.Date] = record
	}
	unavailableDays := make([]time.Time, 0)
	for _, date := range datesToBook {
		if record, found := dateToAvailabilityRecords[date]; !found {
			return model.Order{}, fmt.Errorf("no availability info for date %v", date)
		} else if record.Quota == 0 {
			unavailableDays = append(unavailableDays, date)
		}
	}
	if len(unavailableDays) > 0 {
		return model.Order{}, fmt.Errorf("unavailable dates in specified range: %v", unavailableDays)
	}
	u.orderRepository.Push(order)
	u.availabilityRepository.DecrementQuotaForRange(order.From, order.To)
	u.logger.Infof("order placed: %v", order)
	return order, nil
}

func (u Usecase) validate(order model.Order) error {
	if order.From.IsZero() {
		return errors.New("from date should not be empty")
	}
	if order.From.IsZero() {
		return errors.New("to date should not be empty")
	}
	if order.HotelID == "" {
		return errors.New("hotel id should not be empty")
	}
	if order.RoomID == "" {
		return errors.New("room id should not be empty")
	}
	if order.UserEmail == "" {
		return errors.New("user email should not be empty")
	}
	if order.From.Compare(order.To) != -1 {
		return errors.New("reservation start date should be lesser than end date")
	}
	return nil
}

func getDatesInRange(from time.Time, to time.Time) []time.Time {
	dates := make([]time.Time, 0)
	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates
}

func floorDate(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
