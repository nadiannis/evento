package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/repository"
)

type TicketTypeUsecase struct {
	repository repository.ITicketTypeRepository
}

func NewTicketTypeUsecase(repository repository.ITicketTypeRepository) ITicketTypeUsecase {
	return &TicketTypeUsecase{
		repository: repository,
	}
}

func (u *TicketTypeUsecase) Add(input *request.TicketTypeRequest) (*domain.TicketType, error) {
	ticketType := &domain.TicketType{
		ID:    uuid.NewString(),
		Name:  input.Name,
		Price: input.Price,
	}

	return u.repository.Add(ticketType)
}
