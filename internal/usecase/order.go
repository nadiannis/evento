package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/repository"
)

type OrderUsecase struct {
	orderRepository      repository.IOrderRepository
	customerRepository   repository.ICustomerRepository
	ticketRepository     repository.ITicketRepository
	ticketTypeRepository repository.ITicketTypeRepository
}

func NewOrderUsecase(
	orderRepository repository.IOrderRepository,
	customerRepository repository.ICustomerRepository,
	ticketRepository repository.ITicketRepository,
	ticketTypeRepository repository.ITicketTypeRepository,
) IOrderUsecase {
	return &OrderUsecase{
		orderRepository:      orderRepository,
		customerRepository:   customerRepository,
		ticketRepository:     ticketRepository,
		ticketTypeRepository: ticketTypeRepository,
	}
}

func (u *OrderUsecase) GetAll() []*domain.Order {
	return u.orderRepository.GetAll()
}

func (u *OrderUsecase) Add(input *request.OrderRequest) (*domain.Order, error) {
	customer, err := u.customerRepository.GetByID(input.CustomerID)
	if err != nil {
		return nil, err
	}

	ticket, err := u.ticketRepository.GetByID(input.TicketID)
	if err != nil {
		return nil, err
	}

	ticketType, err := u.ticketTypeRepository.GetByName(ticket.Type)
	if err != nil {
		return nil, err
	}

	order := u.orderRepository.Add(&domain.Order{
		ID:         uuid.NewString(),
		CustomerID: customer.ID,
		TicketID:   ticket.ID,
		Quantity:   input.Quantity,
		TotalPrice: float64(input.Quantity) * ticketType.Price,
		CreatedAt:  time.Now(),
	})

	err = u.ticketRepository.DeductQuantity(ticket.ID, order.Quantity)
	if err != nil {
		return nil, err
	}

	err = u.customerRepository.AddOrder(customer.ID, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
