package repositories

import (
	"LandTicket-Backend/models"
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user models.User) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User)(models.User, error){
	err := r.db.Exec("INSERT INTO users(fullname, username, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?)", user.Fullname, user.Username, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}
