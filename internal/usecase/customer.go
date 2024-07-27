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
		Balance:  input.Balance,
		Orders:   make([]*domain.Order, 0),
	}

	return u.repository.Add(customer)
}

func (u *CustomerUsecase) AddBalance(customerID string, input *request.CustomerBalanceRequest) (*domain.Customer, error) {
	customer, err := u.repository.GetByID(customerID)
	if err != nil {
		return nil, err
	}

	err = u.repository.AddBalance(customerID, input.Balance)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
