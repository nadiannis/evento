package repository

import "github.com/nadiannis/evento/internal/domain"

type EventRepository struct {
	db map[string]*domain.Event
}

func NewEventRepository() IEventRepository {
	return &EventRepository{
		db: make(map[string]*domain.Event),
	}
}

func (r *EventRepository) GetAll() []*domain.Event {
	events := make([]*domain.Event, 0)
	for _, event := range r.db {
		events = append(events, event)
	}
	return events
}

func (r *EventRepository) Add(event *domain.Event) *domain.Event {
	r.db[event.ID] = event
	return event
}
