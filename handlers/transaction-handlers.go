package handlers

import (
	dto "LandTicket-Backend/dto/result"
	"LandTicket-Backend/dto/transaction"
	"LandTicket-Backend/models"
	"LandTicket-Backend/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlersTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepositories repositories.TransactionRepository) *handlersTransaction {
	return &handlersTransaction{TransactionRepositories}
}

func (h *handlersTransaction) CreateTransaction(c echo.Context) error {
	request := new(transaction.CreateTransactionRequest)
	if err:=c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transaction := models.Transaction{
		UserId: request.UserId,
		TicketId: request.TicketId,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *handlersTransaction) FindTransaction(c echo.Context) error {
	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *handlersTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(transaction)})
}


func convertResponseTransaction(t models.Transaction) transaction.TransactionResponse{
	return transaction.TransactionResponse{
		ID: t.ID,
		UserId: t.ID,
		TicketId: t.ID,
	}
}