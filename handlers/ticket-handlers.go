package handlers

import (
	dto "LandTicket-Backend/dto/result"
	ticketdto "LandTicket-Backend/dto/ticket"
	"LandTicket-Backend/models"
	"LandTicket-Backend/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlersTicket struct {
	TicketRepository repositories.TicketRepository
}

func HandlerTicket(TicketRepositories repositories.TicketRepository) *handlersTicket {
	return &handlersTicket{TicketRepositories}
}

func (h *handlersTicket) CreateTicket(c echo.Context) error {
	request := new(ticketdto.CreateTicketRequest)
	if err:=c.Bind(request); err!=nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ticket := models.Ticket{
		NameTrain: request.NameTrain,
		TypeTrain: request.TypeTrain,
		StartDate: request.StartDate,
		StartStationID: request.StartStationID,
		StartTime: request.StartTime,
		DestinationStationID: request.DestinationStationID,
		ArrivalTime: request.ArrivalTime,
		Price: request.Price,
	}

	data, err := h.TicketRepository.CreateTicket(ticket)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTicket(data)})
}

func (h *handlersTicket) FindTicket(c echo.Context) error {
	ticket, err := h.TicketRepository.FindTicket()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: ticket})
}

func (h *handlersTicket) GetTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ticket, err := h.TicketRepository.GetTicket(id)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTicket(ticket)})
}



func convertResponseTicket(t models.Ticket) ticketdto.TicketResponse{
	return ticketdto.TicketResponse{
		ID: t.ID,
		NameTrain: t.NameTrain,
		TypeTrain: t.TypeTrain,
		StartDate: t.StartDate,
		StartStationID: t.StartStationID,
		StartTime: t.StartTime,
		DestinationStationID: t.DestinationStationID,
		ArrivalTime: t.ArrivalTime,
		Price: t.Price,
		Qty: t.Qty,
	}
}