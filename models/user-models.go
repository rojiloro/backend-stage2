package models

import "time"

// user models struct
type User struct {
	ID int			 		`json:"id" gorm:"primarykey:autoIncrement"`
	Fullname string		 	`json:"fullname" gorm:"type: varchar(255)"`
	Username string 		`json:"username" gorm:"type: varchar(255)"`
	Email string 			`json:"email" gorm:"type: varchar(255)"`
	Password string			`json:"-" gorm:"type: varchar(255)"`
	NoHP string `json:"no_hp" gorm:"type : varchar(255)"`
	CreatedAt time.Time		`json:"-"`
	UpdatedAt time.Time		`json:"-"`
}

type UserResponse struct{
	Username string 		`json:"username" form:"username"`
	Email string 			`json:"email" form:"email"`
	NoHP string `json:"no_hp" form:"no_hp"`
}

func (UserResponse) TableName() string {
	return "users"
}