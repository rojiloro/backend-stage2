package routes

import (
	"LandTicket-Backend/handlers"

	"LandTicket-Backend/pkg/mysql"
	"LandTicket-Backend/repositories"

	"github.com/labstack/echo/v4"
)

func TicketRoutes(e *echo.Group){
	TicketRepository := repositories.RepositoryTicket(mysql.DB)
	h := handlers.HandlerTicket(TicketRepository)

	e.POST("/ticket", h.CreateTicket) 
	e.GET("/tickets", h.FindTicket)
	e.GET("/ticket/:id", h.GetTicket )
	e.GET("/ticket", h.FilterTicket)
}