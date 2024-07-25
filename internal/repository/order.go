package repository

import "github.com/nadiannis/evento/internal/domain"

type OrderRepository struct {
	db map[string]*domain.Order
}

func NewOrderRepository() IOrderRepository {
	return &OrderRepository{
		db: make(map[string]*domain.Order),
	}
}

func (r *OrderRepository) GetAll() []*domain.Order {
	orders := make([]*domain.Order, 0)
	for _, order := range r.db {
		orders = append(orders, order)
	}
	return orders
}

func (r *OrderRepository) Add(order *domain.Order) *domain.Order {
	r.db[order.ID] = order
	return order
}
