package handler

import (
	"errors"
	"net/http"

	"github.com/nadiannis/evento/internal/domain/request"
	"github.com/nadiannis/evento/internal/domain/response"
	"github.com/nadiannis/evento/internal/usecase"
	"github.com/nadiannis/evento/internal/utils"
)

type CustomerHandler struct {
	usecase usecase.ICustomerUsecase
}

func NewCustomerHandler(usecase usecase.ICustomerUsecase) ICustomerHandler {
	return &CustomerHandler{
		usecase: usecase,
	}
}

func (h *CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers := h.usecase.GetAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customers retrieved successfully",
		Data:    customers,
	}

	err := utils.WriteJSON(w, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *CustomerHandler) Add(w http.ResponseWriter, r *http.Request) {
	var input request.CustomerRequest

	err := utils.ReadJSON(r, &input)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	customer, err := h.usecase.Add(&input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerAlreadyExists):
			utils.BadRequestResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customer added successfully",
		Data:    customer,
	}

	err = utils.WriteJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
