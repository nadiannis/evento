package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/repository"
)

type TicketUsecase struct {
	ticketRepository repository.ITicketRepository
	eventRepository  repository.IEventRepository
}

func NewTicketUsecase(ticketRepository repository.ITicketRepository, eventRepository repository.IEventRepository) ITicketUsecase {
	return &TicketUsecase{
		ticketRepository: ticketRepository,
		eventRepository:  eventRepository,
	}
}

func (u *TicketUsecase) GetAll() []*domain.Ticket {
	return u.ticketRepository.GetAll()
}

func (u *TicketUsecase) Add(input *request.TicketRequest) (*domain.Ticket, error) {
	ticket := &domain.Ticket{
		ID:       uuid.NewString(),
		EventID:  input.EventID,
		Type:     input.Type,
		Quantity: input.Quantity,
	}

	savedTicket := u.ticketRepository.Add(ticket)
	_, err := u.eventRepository.AddTicket(savedTicket.EventID, savedTicket)
	if err != nil {
		return nil, err
	}

	return savedTicket, nil
}
