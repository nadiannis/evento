package usecase

import (
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
)

type CustomerReader interface {
	GetAll() []*domain.Customer
}

type CustomerWriter interface {
	Add(input *request.CustomerRequest) (*domain.Customer, error)
	AddBalance(customerID string, input *request.CustomerBalanceRequest) (*domain.Customer, error)
}

type ICustomerUsecase interface {
	CustomerReader
	CustomerWriter
}

type EventReader interface {
	GetAll() []*domain.Event
	GetByID(eventID string) (*domain.Event, error)
}

type EventWriter interface {
	Add(input *request.EventRequest) *domain.Event
}

type IEventUsecase interface {
	EventReader
	EventWriter
}

type TicketTypeWriter interface {
	Add(input *request.TicketTypeRequest) (*domain.TicketType, error)
}

type ITicketTypeUsecase interface {
	TicketTypeWriter
}

type TicketReader interface {
	GetAll() []*domain.Ticket
	GetByID(ticketID string) (*domain.Ticket, error)
}

type TicketWriter interface {
	Add(input *request.TicketRequest) (*domain.Ticket, error)
}

type ITicketUsecase interface {
	TicketReader
	TicketWriter
}

type OrderReader interface {
	GetAll() []*domain.Order
}

type OrderWriter interface {
	Add(input *request.OrderRequest) (*domain.Order, error)
}

type IOrderUsecase interface {
	OrderReader
	OrderWriter
}
