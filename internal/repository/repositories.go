package repository

type Repositories struct {
	Customers   ICustomerRepository
	Events      IEventRepository
	TicketTypes ITicketTypeRepository
	Tickets     ITicketRepository
	Orders      IOrderRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Customers:   NewCustomerRepository(),
		Events:      NewEventRepository(),
		TicketTypes: NewTicketTypeRepository(),
		Tickets:     NewTicketRepository(),
		Orders:      NewOrderRepository(),
	}
}
