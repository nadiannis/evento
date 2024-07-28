package repository

import (
	"sync"

	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/utils"
)

type CustomerRepository struct {
	db map[string]*domain.Customer
	mu sync.Mutex
}

func NewCustomerRepository() ICustomerRepository {
	return &CustomerRepository{
		db: make(map[string]*domain.Customer),
	}
}

func (r *CustomerRepository) GetAll() []*domain.Customer {
	r.mu.Lock()
	defer r.mu.Unlock()

	customers := make([]*domain.Customer, 0)
	for _, customer := range r.db {
		customers = append(customers, customer)
	}
	return customers
}

func (r *CustomerRepository) Add(customer *domain.Customer) (*domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, c := range r.db {
		if c.Username == customer.Username {
			return nil, utils.ErrCustomerAlreadyExists
		}
	}

	r.db[customer.ID] = customer
	return customer, nil
}

func (r *CustomerRepository) GetByID(customerID string) (*domain.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if customer, exists := r.db[customerID]; exists {
		return customer, nil
	}

	return nil, utils.ErrCustomerNotFound
}

func (r *CustomerRepository) AddOrder(customerID string, order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if customer, exists := r.db[customerID]; exists {
		customer.Orders = append(customer.Orders, order)
		return nil
	}

	return utils.ErrCustomerNotFound
}

func (r *CustomerRepository) DeleteAllOrders() {
	for _, customer := range r.db {
		customer.Orders = make([]*domain.Order, 0)
	}
}

func (r *CustomerRepository) AddBalance(customerID string, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	customer, exists := r.db[customerID]
	if !exists {
		return utils.ErrCustomerNotFound
	}

	customer.Balance += amount
	r.db[customerID] = customer
	return nil
}

func (r *CustomerRepository) DeductBalance(customerID string, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	customer, exists := r.db[customerID]
	if !exists {
		return utils.ErrCustomerNotFound
	}

	if customer.Balance < amount {
		return utils.ErrInsufficientBalance
	}

	customer.Balance -= amount
	r.db[customerID] = customer
	return nil
}
