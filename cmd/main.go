package main

import (
	"fmt"
	"net/http"

	"github.com/nadiannis/evento/internal/handler"
	"github.com/nadiannis/evento/internal/repository"
	"github.com/nadiannis/evento/internal/usecase"
	"github.com/rs/zerolog/log"
)

type application struct {
	port     int
	handlers handler.Handlers
}

func main() {
	repos := repository.NewRepositories()
	usecases := usecase.NewUsecases(repos)
	handlers := handler.NewHandlers(usecases)

	app := &application{
		port:     8080,
		handlers: handlers,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.port),
		Handler: app.routes(),
	}

	log.Info().Msg("starting server on " + srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
