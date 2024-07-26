package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/nadiannis/evento/internal/domain/response"
)

func ReadJSON(r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		default:
			return err
		}
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data response.Response, headers http.Header) error {
	jsonBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	jsonBytes = append(jsonBytes, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)

	var message any
	switch v := data.(type) {
	case response.SuccessResponse:
		message = v.Message
	case response.ErrorResponse:
		message = v.Detail
	default:
		message = "request processed"
	}

	*r = *r.WithContext(SetLogMessage(r.Context(), message))

	return nil
}
