package repository

type Repositories struct {
	Customers   ICustomerRepository
	Events      IEventRepository
	TicketTypes ITicketTypeRepository
	Tickets     ITicketRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Customers:   NewCustomerRepository(),
		Events:      NewEventRepository(),
		TicketTypes: NewTicketTypeRepository(),
		Tickets:     NewTicketRepository(),
	}
}
