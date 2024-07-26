package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/nadiannis/evento/internal/domain/response"
	"github.com/rs/zerolog/log"
)

var (
	ErrCustomerNotFound           = errors.New("customer not found")
	ErrTicketTypeNotFound         = errors.New("ticket type not found")
	ErrTicketNotFound             = errors.New("ticket not found")
	ErrEventNotFound              = errors.New("event not found")
	ErrCustomerAlreadyExists      = errors.New("customer already exists")
	ErrTicketTypeAlreadyExists    = errors.New("ticket type already exists")
	ErrTicketAlreadyExists        = errors.New("ticket already exists for the event")
	ErrInsufficientTicketQuantity = errors.New("insufficient ticket quantity")
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	res := response.ErrorResponse{
		Status:  "error",
		Message: strings.ToLower(http.StatusText(status)),
		Detail:  message,
	}

	err := WriteJSON(w, status, res, nil)
	if err != nil {
		req := fmt.Sprint(r.Method, " ", r.URL.String())
		log.Error().Str("request", req).Err(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg(err.Error())

	message := "server encountered a problem"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg(err.Error())

	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg(err.Error())

	errorResponse(w, r, http.StatusNotFound, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	req := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
	log.Error().Str("request", req).Msg("invalid request body")

	errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
