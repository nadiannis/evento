package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/repository"
)

type CustomerUsecase struct {
	repository repository.ICustomerRepository
}

func NewCustomerUsecase(repository repository.ICustomerRepository) ICustomerUsecase {
	return &CustomerUsecase{
		repository: repository,
	}
}

func (u *CustomerUsecase) GetAll() []*domain.Customer {
	return u.repository.GetAll()
}

func (u *CustomerUsecase) Add(input *request.CustomerRequest) (*domain.Customer, error) {
	customer := &domain.Customer{
		ID:       uuid.NewString(),
		Username: input.Username,
		Orders:   make([]*domain.Order, 0),
	}

	return u.repository.Add(customer)
}
