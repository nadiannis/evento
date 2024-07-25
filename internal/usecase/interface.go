package usecase

import (
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
)

type ICustomerUsecase interface {
	GetAll() []*domain.Customer
	Add(input *request.CustomerRequest) (*domain.Customer, error)
}

type IEventUsecase interface {
	GetAll() []*domain.Event
	Add(input *request.EventRequest) *domain.Event
	GetByID(eventID string) (*domain.Event, error)
}

type ITicketTypeUsecase interface {
	Add(input *request.TicketTypeRequest) (*domain.TicketType, error)
}

type ITicketUsecase interface {
	Add(input *request.TicketRequest) (*domain.Ticket, error)
}

type IOrderUsecase interface {
	GetAll() []*domain.Order
}
