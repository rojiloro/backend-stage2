package repositories

import (
	"LandTicket-Backend/models"
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) (models.User, error)
}

type authrepository struct {
	db *gorm.DB
}

func (r *repository) CreateUser (user models.User) (models.User, error){
	err := r.db.Exec("INSERT INTO users(fullname, username, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?)", user.Fullname, user.Username, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}
