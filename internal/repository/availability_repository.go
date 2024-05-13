package repository

import (
	"applicationDesignTest/internal/model"
	"errors"
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

func (r availabilityRepository) GetByIndex(index int) (model.RoomAvailability, error) {
	if index < 0 {
		return model.RoomAvailability{}, errors.New("invalid index")
	}
	if index >= len(r.availabilityRecords) {
		return model.RoomAvailability{}, errors.New("index out of range")
	}
	return r.availabilityRecords[index], nil
}

func (r availabilityRepository) GetAll() []model.RoomAvailability {
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

func (r *availabilityRepository) Close() {
	r.availabilityRecords = []model.RoomAvailability{}
}
