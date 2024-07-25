package repository

import "github.com/nadiannis/evento/internal/domain"

type ICustomerRepository interface {
	GetAll() []*domain.Customer
	Add(user *domain.Customer) (*domain.Customer, error)
}

type IEventRepository interface {
	GetAll() []*domain.Event
	Add(event *domain.Event) *domain.Event
}
