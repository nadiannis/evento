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

	mux.HandleFunc("GET /api/events", app.handlers.Events.GetAll)
	mux.HandleFunc("GET /api/events/{id}", app.handlers.Events.GetByID)

	mux.HandleFunc("GET /api/orders", app.handlers.Orders.GetAll)

	return requestLogger(mux)
}
