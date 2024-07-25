package repository

type Repositories struct {
	Customers ICustomerRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Customers: NewCustomerRepository(),
	}
}
