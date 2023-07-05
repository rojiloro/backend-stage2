package mysql

import (
	"LandTicket-Backend/pkg/mysql"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

// connection Database
func DatabaseInit(){
	var err error
	dsn := "root:@tcp(localhost:3306)/land-tick?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database berhasil ges!!!")
}