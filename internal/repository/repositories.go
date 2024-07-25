package repository

type Repositories struct {
	Customers   ICustomerRepository
	Events      IEventRepository
	TicketTypes ITicketTypeRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Customers:   NewCustomerRepository(),
		Events:      NewEventRepository(),
		TicketTypes: NewTicketTypeRepository(),
	}
}
