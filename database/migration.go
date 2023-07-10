package database

import (
	"LandTicket-Backend/models"
	"LandTicket-Backend/pkg/mysql"
	"fmt"
)

// automatic Migration if running app
func RunMigration(){
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Station{},
		&models.Ticket{},
		&models.Transaction{},
	)

	if err != nil {
		panic("migration failed")
	}

	fmt.Println("Migration success gess!!!")
}