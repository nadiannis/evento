package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
		log.Info().Msg(request)

		next.ServeHTTP(w, r)
	})
}
