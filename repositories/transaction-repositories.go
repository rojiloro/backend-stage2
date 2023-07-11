package repositories

import (
	"LandTicket-Backend/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction)(models.Transaction, error)
	FindTransaction()([]models.Transaction, error)
	GetTransaction(ID int)(models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction)(models.Transaction, error){
	err := r.db.Preload("user").Preload("ticket").Create(&transaction).Error

	return transaction, err
}

func (r *repository) FindTransaction()([]models.Transaction, error){
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int)(models.Transaction, error){
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Preload("Ticket.StartStation").Preload("Ticket.DestinationStation").First(&transaction, ID).Error

	return transaction, err
}