package usecase

import "github.com/nadiannis/evento/internal/repository"

type Usecases struct {
	Customers   ICustomerUsecase
	Events      IEventUsecase
	TicketTypes ITicketTypeUsecase
	Tickets     ITicketUsecase
	Orders      IOrderUsecase
}

func NewUsecases(repositories repository.Repositories) Usecases {
	return Usecases{
		Customers:   NewCustomerUsecase(repositories.Customers),
		Events:      NewEventUsecase(repositories.Events),
		TicketTypes: NewTicketTypeUsecase(repositories.TicketTypes),
		Tickets:     NewTicketUsecase(repositories.Tickets, repositories.Events),
		Orders:      NewOrderUsecase(repositories.Orders),
	}
}
