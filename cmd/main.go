package main

import (
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

	log.Info().Msg("add ticket types")
	prepopulateTicketTypes(usecases.TicketTypes)

	log.Info().Msg("add events and tickets")
	prepopulateEventsAndTickets(usecases.Events, usecases.Tickets)

	err := app.serve()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
