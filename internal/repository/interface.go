package repository

import "github.com/nadiannis/evento/internal/domain"

type CustomerReader interface {
	GetAll() []*domain.Customer
	GetByID(customerID string) (*domain.Customer, error)
}

type CustomerWriter interface {
	Add(customer *domain.Customer) (*domain.Customer, error)
	AddBalance(customerID string, amount float64) error
	DeductBalance(customerID string, amount float64) error
	AddOrder(customerID string, order *domain.Order) error
	DeleteAllOrders()
}

type ICustomerRepository interface {
	CustomerReader
	CustomerWriter
}

type EventReader interface {
	GetAll() []*domain.Event
	GetByID(eventID string) (*domain.Event, error)
}

type EventWriter interface {
	Add(event *domain.Event) *domain.Event
	AddTicket(eventID string, ticket *domain.Ticket) (*domain.Ticket, error)
}

type IEventRepository interface {
	EventReader
	EventWriter
}

type TicketTypeReader interface {
	GetByName(ticketTypeName domain.TicketTypeName) (*domain.TicketType, error)
}

type TicketTypeWriter interface {
	Add(ticketType *domain.TicketType) (*domain.TicketType, error)
}

type ITicketTypeRepository interface {
	TicketTypeReader
	TicketTypeWriter
}

type TicketReader interface {
	GetAll() []*domain.Ticket
	GetByID(ticketID string) (*domain.Ticket, error)
}

type TicketWriter interface {
	Add(ticket *domain.Ticket) (*domain.Ticket, error)
	AddQuantity(ticketID string, quantity int) error
	DeductQuantity(ticketID string, quantity int) error
}

type ITicketRepository interface {
	TicketReader
	TicketWriter
}

type OrderReader interface {
	GetAll() []*domain.Order
}

type OrderWriter interface {
	Add(order *domain.Order) *domain.Order
	DeleteByID(orderID string) error
	DeleteAll()
}

type IOrderRepository interface {
	OrderReader
	OrderWriter
}
