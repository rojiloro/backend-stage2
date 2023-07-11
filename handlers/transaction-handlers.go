package handlers

import (
	dto "LandTicket-Backend/dto/result"
	"LandTicket-Backend/dto/transaction"
	"LandTicket-Backend/models"
	"LandTicket-Backend/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlersTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepositories repositories.TransactionRepository) *handlersTransaction {
	return &handlersTransaction{TransactionRepositories}
}

func (h *handlersTransaction) CreateTransaction(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	fmt.Println(user_id)
	ticket_id, _ := strconv.Atoi(c.FormValue("ticket_id"))
	fmt.Println(ticket_id)

	request := transaction.CreateTransactionRequest{
		TicketId: ticket_id,
		Image : dataFile,
		Status: c.FormValue("status"),
		UserId: user_id,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transaction := models.Transaction{
		UserId: request.UserId,
		TicketId: request.TicketId,
		Image: request.Image,
		Status: request.Status,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlersTransaction) FindTransaction(c echo.Context) error {
	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	
	for i, p := range transaction {
		transaction[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *handlersTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}