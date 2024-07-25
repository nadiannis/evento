package handler

import (
	"net/http"

	"github.com/nadiannis/evento/internal/domain/response"
	"github.com/nadiannis/evento/internal/usecase"
	"github.com/nadiannis/evento/internal/utils"
)

type OrderHandler struct {
	usecase usecase.IOrderUsecase
}

func NewOrderHandler(usecase usecase.IOrderUsecase) IOrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orders := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "orders retrieved successfully",
		Data:    orders,
	}

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
