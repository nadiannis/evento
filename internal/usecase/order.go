package usecase

import (
	"github.com/nadiannis/evento/internal/domain"
	"github.com/nadiannis/evento/internal/repository"
)

type OrderUsecase struct {
	repository repository.IOrderRepository
}

func NewOrderUsecase(repository repository.IOrderRepository) IOrderUsecase {
	return &OrderUsecase{
		repository: repository,
	}
}

func (u *OrderUsecase) GetAll() []*domain.Order {
	return u.repository.GetAll()
}
