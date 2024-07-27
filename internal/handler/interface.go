package handler

import "net/http"

type CustomerReader interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type CustomerWriter interface {
	Add(w http.ResponseWriter, r *http.Request)
	AddBalance(w http.ResponseWriter, r *http.Request)
}

type ICustomerHandler interface {
	CustomerReader
	CustomerWriter
}

type EventReader interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type IEventHandler interface {
	EventReader
}

type TicketReader interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type ITicketHandler interface {
	TicketReader
}

type OrderReader interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type OrderWriter interface {
	Add(w http.ResponseWriter, r *http.Request)
}

type IOrderHandler interface {
	OrderReader
	OrderWriter
}
