package main

import (
	"fmt"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API is running on port %d\n", app.port)
	})

	mux.HandleFunc("GET /api/customers", app.handlers.Customers.GetAll)
	mux.HandleFunc("POST /api/customers", app.handlers.Customers.Add)
	mux.HandleFunc("GET /api/customers/{id}", app.handlers.Customers.GetByID)
	mux.HandleFunc("PATCH /api/customers/{id}/balances", app.handlers.Customers.AddBalance)

	mux.HandleFunc("GET /api/events", app.handlers.Events.GetAll)
	mux.HandleFunc("GET /api/events/{id}", app.handlers.Events.GetByID)

	mux.HandleFunc("GET /api/tickets", app.handlers.Tickets.GetAll)
	mux.HandleFunc("GET /api/tickets/{id}", app.handlers.Tickets.GetByID)
	mux.HandleFunc("PATCH /api/tickets/{id}/quantities", app.handlers.Tickets.AddQuantity) // Intended solely for concurrency testing purpose

	mux.HandleFunc("GET /api/orders", app.handlers.Orders.GetAll)
	mux.HandleFunc("POST /api/orders", app.handlers.Orders.Add)
	mux.HandleFunc("DELETE /api/orders", app.handlers.Orders.DeleteAll) // Intended solely for concurrency testing purpose

	return requestLogger(mux)
}
