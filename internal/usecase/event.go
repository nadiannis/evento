package usecase

import (
	"github.com/nadiannis/evento/internal/domain"
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
