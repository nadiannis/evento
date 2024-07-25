package repository

import "github.com/nadiannis/evento/internal/domain"

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
