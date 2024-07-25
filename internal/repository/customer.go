package repository

import (
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/utils"
)

type CustomerRepository struct {
	db map[string]*domain.Customer
}

func NewCustomerRepository() ICustomerRepository {
	return &CustomerRepository{
		db: make(map[string]*domain.Customer),
	}
}

func (r *CustomerRepository) GetAll() []*domain.Customer {
	customers := make([]*domain.Customer, 0)
	for _, customer := range r.db {
		customers = append(customers, customer)
	}
	return customers
}

func (r *CustomerRepository) Add(customer *domain.Customer) (*domain.Customer, error) {
	for _, c := range r.db {
		if c.Username == customer.Username {
			return nil, utils.ErrCustomerAlreadyExists
		}
	}

	r.db[customer.ID] = customer
	return customer, nil
}

func (r *CustomerRepository) GetByID(customerID string) (*domain.Customer, error) {
	if customer, exists := r.db[customerID]; exists {
		return customer, nil
	}

	return nil, utils.ErrCustomerNotFound
}

func (r *CustomerRepository) AddOrder(customerID string, order *domain.Order) error {
	if customer, exists := r.db[customerID]; exists {
		customer.Orders = append(customer.Orders, order)
		return nil
	}

	return utils.ErrCustomerNotFound
}
