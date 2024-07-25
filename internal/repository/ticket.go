package repository

import (
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/utils"
)

type TicketRepository struct {
	db map[string]*domain.Ticket
}

func NewTicketRepository() ITicketRepository {
	return &TicketRepository{
		db: make(map[string]*domain.Ticket),
	}
}

func (r *TicketRepository) Add(ticket *domain.Ticket) *domain.Ticket {
	r.db[ticket.ID] = ticket
	return ticket
}

func (r *TicketRepository) GetByID(ticketID string) (*domain.Ticket, error) {
	if ticket, exists := r.db[ticketID]; exists {
		return ticket, nil
	}

	return nil, utils.ErrTicketNotFound
}

func (r *TicketRepository) DeductQuantity(ticketID string, quantity int) error {
	ticket, exists := r.db[ticketID]
	if !exists {
		return utils.ErrTicketNotFound
	}

	if ticket.Quantity < quantity {
		return utils.ErrInsufficientTicketQuantity
	}

	ticket.Quantity -= quantity
	r.db[ticketID] = ticket
	return nil
}
