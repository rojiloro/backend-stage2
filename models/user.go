package models

import "time"

// user models struct
type User struct {
	ID int			 		`json : "id"`
	Fullname string		 	`json : "fullname" gorm: "type: varchar(255)"`
	Username string 		`json : "username" gorm: "type: varchar(255)"`
	Email string 			`json: "email" gorm: "type: "varchar(255)"`
	Password string			`json: "password" gorm: "type" : "password"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
}