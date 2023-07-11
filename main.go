package main

import (
	"LandTicket-Backend/database"
	"LandTicket-Backend/pkg/mysql"
	"LandTicket-Backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	e.Logger.Fatal(e.Start("localhost:5000"))
}