package repository

import (
	"sync"

	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/utils"
)

type EventRepository struct {
	db map[string]*domain.Event
	mu sync.Mutex
}

func NewEventRepository() IEventRepository {
	return &EventRepository{
		db: make(map[string]*domain.Event),
	}
}

func (r *EventRepository) GetAll() []*domain.Event {
	r.mu.Lock()
	defer r.mu.Unlock()

	events := make([]*domain.Event, 0)
	for _, event := range r.db {
		events = append(events, event)
	}
	return events
}

func (r *EventRepository) Add(event *domain.Event) *domain.Event {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.db[event.ID] = event
	return event
}

func (r *EventRepository) GetByID(eventID string) (*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if event, exists := r.db[eventID]; exists {
		return event, nil
	}

	return nil, utils.ErrEventNotFound
}

func (r *EventRepository) AddTicket(eventID string, ticket *domain.Ticket) (*domain.Ticket, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if event, exists := r.db[eventID]; exists {
		event.Tickets[ticket.Type] = ticket
		return ticket, nil
	}

	return nil, utils.ErrEventNotFound
}
