package handler

import (
	"errors"
	"net/http"
	"time"

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

	time.Sleep(2 * time.Second) // Simulate real processing time

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customers retrieved successfully",
		Data:    customers,
	}

	err := utils.WriteJSON(w, r, http.StatusOK, res, nil)
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

	v := utils.NewValidator()

	v.Check(input.Username != "", "username", "username is required")
	v.Check(input.Balance >= 0, "balance", "balance should not be a negative number")

	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
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

	err = utils.WriteJSON(w, r, http.StatusCreated, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *CustomerHandler) AddBalance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var input request.CustomerBalanceRequest

	err := utils.ReadJSON(r, &input)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.Balance != 0, "balance", "balance is required")
	v.Check(input.Balance > 0, "balance", "balance should not be a negative number")

	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
		return
	}

	customer, err := h.usecase.AddBalance(id, &input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerNotFound):
			utils.NotFoundResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customer balance added successfully",
		Data:    customer,
	}

	err = utils.WriteJSON(w, r, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}

func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	customer, err := h.usecase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerNotFound):
			utils.NotFoundResponse(w, r, err)
		default:
			utils.ServerErrorResponse(w, r, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customer retrieved successfully",
		Data:    customer,
	}

	err = utils.WriteJSON(w, r, http.StatusOK, res, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
