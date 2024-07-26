package handler

import (
	"errors"
	"net/http"

	"github.com/nadiannis/evento/internal/domain/response"
	"github.com/nadiannis/evento/internal/usecase"
	"github.com/nadiannis/evento/internal/utils"
)

type TicketHandler struct {
	usecase usecase.ITicketUsecase
}

func NewTicketHandler(usecase usecase.ITicketUsecase) ITicketHandler {
	return &TicketHandler{
		usecase: usecase,
	}
}

func (h *TicketHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tickets := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "tickets retrieved successfully",
		Data:    tickets,
	}

	err := utils.WriteJSON(w, r, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *TicketHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ticket, err := h.usecase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrTicketNotFound):
			utils.NotFoundResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "ticket retrieved successfully",
		Data:    ticket,
	}

	err = utils.WriteJSON(w, r, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
