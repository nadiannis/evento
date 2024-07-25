package usecase

import "github.com/nadiannis/evento/internal/repository"

type Usecases struct {
	Customers ICustomerUsecase
}

func NewUsecases(repositories repository.Repositories) Usecases {
	return Usecases{
		Customers: NewCustomerUsecase(repositories.Customers),
	}
}
