package handler

import "github.com/nadiannis/evento/internal/usecase"

type Handlers struct {
	Customers ICustomerHandler
	Events    IEventHandler
	Orders    IOrderHandler
}

func NewHandlers(usecases usecase.Usecases) Handlers {
	return Handlers{
		Customers: NewCustomerHandler(usecases.Customers),
		Events:    NewEventHandler(usecases.Events),
		Orders:    NewOrderHandler(usecases.Orders),
	}
}
