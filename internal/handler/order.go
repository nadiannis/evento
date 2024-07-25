package handler

import (
	"errors"
	"net/http"

	"github.com/nadiannis/evento/internal/domain/request"
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

func (h *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	var input request.OrderRequest

	err := utils.ReadJSON(r, &input)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.CustomerID != "", "customer_id", "customer_id is required")
	v.Check(input.TicketID != "", "ticket_id", "ticket_id is required")
	v.Check(input.Quantity > 0, "quantity", "quantity should be greater than 0")

	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
		return
	}

	order, err := h.usecase.Add(&input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerNotFound) || errors.Is(err, utils.ErrTicketNotFound) || errors.Is(err, utils.ErrTicketTypeNotFound):
			utils.NotFoundResponse(w, r, err)
		case errors.Is(err, utils.ErrInsufficientTicketQuantity):
			utils.BadRequestResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "order added successfully",
		Data:    order,
	}

	err = utils.WriteJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
