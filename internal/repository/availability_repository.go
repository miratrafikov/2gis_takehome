package repository

import (
	"applicationDesignTest/internal/model"
	"errors"
	"time"
)

type availabilityRepository struct {
	availabilityRecords []model.RoomAvailability
}

func NewAvailabilityRepository() *availabilityRepository {
	return &availabilityRepository{}
}

func (r *availabilityRepository) Push(order model.RoomAvailability) []model.RoomAvailability {
	r.availabilityRecords = append(r.availabilityRecords, order)
	return r.availabilityRecords
}

func (r *availabilityRepository) GetByIndex(index int) (model.RoomAvailability, error) {
	if index < 0 {
		return model.RoomAvailability{}, errors.New("invalid index")
	}
	if index >= len(r.availabilityRecords) {
		return model.RoomAvailability{}, errors.New("index out of range")
	}
	return r.availabilityRecords[index], nil
}

func (r *availabilityRepository) GetAll() []model.RoomAvailability {
	return r.availabilityRecords
}

func (r *availabilityRepository) Replace(index int, record model.RoomAvailability) error {
	if index < 0 {
		return errors.New("invalid index")
	}
	if index >= len(r.availabilityRecords) {
		return errors.New("index out of range")
	}
	r.availabilityRecords[index] = record
	return nil
}

func (r *availabilityRepository) DecrementQuotaForRange(from, to time.Time) {
	// assume map is used or SQL query is used that updates records in range...
	// for d := from; d <= to; d = d.AddDate(0, 0, 1) {
	// 	record := r.availabilityRecords[d]
	// 	record.Quota--
	// 	r.availabilityRecords[d] = record
	// }
}

func (r *availabilityRepository) Close() {
	r.availabilityRecords = []model.RoomAvailability{}
}

func (r *availabilityRepository) LoadFixtures() {
	date := func(year, month, day int) time.Time {
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}
	availabilityRecords := []model.RoomAvailability{
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 1), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 2), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 3), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 4), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 5), Quota: 0},
	}
	for i := 0; i < len(availabilityRecords); i++ {
		r.Push(availabilityRecords[i])
	}
}
