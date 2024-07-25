package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/repository"
)

type EventUsecase struct {
	repository repository.IEventRepository
}

func NewEventUsecase(repository repository.IEventRepository) IEventUsecase {
	return &EventUsecase{
		repository: repository,
	}
}

func (u *EventUsecase) GetAll() []*domain.Event {
	return u.repository.GetAll()
}

func (u *EventUsecase) Add(input *request.EventRequest) *domain.Event {
	event := &domain.Event{
		ID:      uuid.NewString(),
		Name:    input.Name,
		Date:    input.Date,
		Tickets: make(map[domain.TicketTypeName]*domain.Ticket),
	}

	return u.repository.Add(event)
}
