package repository

import (
	"applicationDesignTest/internal/model"
	"errors"
)

type OrderRepository struct {
	orders []model.Order
}

func (r *OrderRepository) Push(order model.Order) []model.Order {
	r.orders = append(r.orders, order)
	return r.orders
}

func (r OrderRepository) GetByIndex(index int) (model.Order, error) {
	if index < 0 {
		return model.Order{}, errors.New("invalid index")
	}
	if index >= len(r.orders) {
		return model.Order{}, errors.New("index out of range")
	}
	return r.orders[index], nil
}
