package repository

import (
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/utils"
)

type TicketTypeRepository struct {
	db map[string]*domain.TicketType
}

func NewTicketTypeRepository() ITicketTypeRepository {
	return &TicketTypeRepository{
		db: make(map[string]*domain.TicketType),
	}
}

func (r *TicketTypeRepository) Add(ticketType *domain.TicketType) (*domain.TicketType, error) {
	for _, t := range r.db {
		if t.Name == ticketType.Name {
			return nil, utils.ErrTicketTypeAlreadyExists
		}
	}

	r.db[ticketType.ID] = ticketType
	return ticketType, nil
}
