package repository

import (
	"sync"

	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/utils"
)

type TicketTypeRepository struct {
	db map[string]*domain.TicketType
	mu sync.Mutex
}

func NewTicketTypeRepository() ITicketTypeRepository {
	return &TicketTypeRepository{
		db: make(map[string]*domain.TicketType),
	}
}

func (r *TicketTypeRepository) Add(ticketType *domain.TicketType) (*domain.TicketType, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range r.db {
		if t.Name == ticketType.Name {
			return nil, utils.ErrTicketTypeAlreadyExists
		}
	}

	r.db[ticketType.ID] = ticketType
	return ticketType, nil
}

func (r *TicketTypeRepository) GetByName(ticketTypeName domain.TicketTypeName) (*domain.TicketType, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, ticketType := range r.db {
		if ticketType.Name == ticketTypeName {
			return ticketType, nil
		}
	}

	return nil, utils.ErrTicketTypeNotFound
}
