package main

import (
	"fmt"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API is running on port %d", app.port)
	})

	mux.HandleFunc("GET /api/customers", app.handlers.Customers.GetAll)
	mux.HandleFunc("POST /api/customers", app.handlers.Customers.Add)

	return requestLogger(mux)
}
