package repository

type Repositories struct {
	Customers ICustomerRepository
	Events    IEventRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Customers: NewCustomerRepository(),
		Events:    NewEventRepository(),
	}
}
