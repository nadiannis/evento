package handler

import "net/http"

type ICustomerHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type IEventHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}
