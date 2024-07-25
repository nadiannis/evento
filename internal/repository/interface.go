package repository

import "github.com/nadiannis/evento/internal/domain"

type ICustomerRepository interface {
	GetAll() []*domain.Customer
	Add(customer *domain.Customer) (*domain.Customer, error)
	GetByID(customerID string) (*domain.Customer, error)
	AddOrder(customerID string, order *domain.Order) error
}

type IEventRepository interface {
	GetAll() []*domain.Event
	Add(event *domain.Event) *domain.Event
	GetByID(eventID string) (*domain.Event, error)
	AddTicket(eventID string, ticket *domain.Ticket) (*domain.Ticket, error)
}

type ITicketTypeRepository interface {
	Add(ticketType *domain.TicketType) (*domain.TicketType, error)
	GetByName(ticketTypeName domain.TicketTypeName) (*domain.TicketType, error)
}

type ITicketRepository interface {
	Add(ticket *domain.Ticket) *domain.Ticket
	GetByID(ticketID string) (*domain.Ticket, error)
	DeductQuantity(ticketID string, quantity int) error
}

type IOrderRepository interface {
	GetAll() []*domain.Order
	Add(order *domain.Order) *domain.Order
}
