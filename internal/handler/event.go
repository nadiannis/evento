package handler

import (
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

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
