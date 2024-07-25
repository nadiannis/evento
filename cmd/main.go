package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type application struct {
	port int
}

func main() {
	app := &application{
		port: 8080,
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
