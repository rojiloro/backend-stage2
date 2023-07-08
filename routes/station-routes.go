package routes

import (
	"LandTicket-Backend/handlers"
	"LandTicket-Backend/pkg/mysql"
	"LandTicket-Backend/repositories"

	"github.com/labstack/echo/v4"
)

func StationRoutes(e *echo.Group){
	StationRepository := repositories.RepositoryStation(mysql.DB)
	h := handlers.HandlerStation(StationRepository)

	e.POST("/station", h.CreateStation)
	e.GET("/station", h.FindStation)
}