package handler

import (
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

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
