package handler

import (
	"errors"
	"net/http"

	"github.com/nadiannis/evento/internal/domain/response"
	"github.com/nadiannis/evento/internal/usecase"
	"github.com/nadiannis/evento/internal/utils"
)

type EventHandler struct {
	usecase usecase.IEventUsecase
}

func NewEventHandler(usecase usecase.IEventUsecase) IEventHandler {
	return &EventHandler{
		usecase: usecase,
	}
}

func (h *EventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	events := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "events retrieved successfully",
		Data:    events,
	}

	err := utils.WriteJSON(w, r, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *EventHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	event, err := h.usecase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrEventNotFound):
			utils.NotFoundResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "event retrieved successfully",
		Data:    event,
	}

	err = utils.WriteJSON(w, r, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
