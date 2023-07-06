package repositories

import (
	"LandTicket-Backend/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUser() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser() ([]models.User, error){
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error){
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) CreateUser (user models.User) (models.User, error){
	err := r.db.Exec("INSERT INTO users(fullname, username, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?)", user.Fullname, user.Username, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}

func (r *repository) UpdateUser (user models.User, ID int)(models.User, error){
	err := r.db.Raw("UPDATE users SET fullname=?, username=?, email=?, password=? WHERE id=?", user.Fullname, user.Username, user.Email, user.Password,ID).Scan(&user).Error

	return user, err
}