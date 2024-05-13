package repository

import "applicationDesignTest/internal/model"

type orderRepository struct {
	orders []model.Order
}

func NewOrderRepository() *orderRepository {
	return &orderRepository{}
}

func (r *orderRepository) Push(order model.Order) []model.Order {
	r.orders = append(r.orders, order)
	return r.orders
}

func (r *orderRepository) Close() {
	r.orders = []model.Order{}
}
