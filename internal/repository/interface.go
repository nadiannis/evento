package repository

import "github.com/nadiannis/evento/internal/domain"

type ICustomerRepository interface {
	GetAll() []*domain.Customer
	Add(user *domain.Customer) (*domain.Customer, error)
}

type IEventRepository interface {
	GetAll() []*domain.Event
	Add(event *domain.Event) *domain.Event
	GetByID(eventID string) (*domain.Event, error)
	AddTicket(eventID string, ticket *domain.Ticket) (*domain.Ticket, error)
}

type ITicketTypeRepository interface {
	Add(ticketType *domain.TicketType) (*domain.TicketType, error)
}

type ITicketRepository interface {
	Add(ticket *domain.Ticket) *domain.Ticket
}

type IOrderRepository interface {
	GetAll() []*domain.Order
}
