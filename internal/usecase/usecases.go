package usecase

import "github.com/nadiannis/evento/internal/repository"

type Usecases struct {
	Customers ICustomerUsecase
	Events    IEventUsecase
}

func NewUsecases(repositories repository.Repositories) Usecases {
	return Usecases{
		Customers: NewCustomerUsecase(repositories.Customers),
		Events:    NewEventUsecase(repositories.Events),
	}
}
