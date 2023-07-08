package repositories

import (
	"LandTicket-Backend/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	CreateTicket(ticket models.Ticket)(models.Ticket, error)
	FindTicket()([]models.Ticket, error)
	GetTicket(ID int) (models.Ticket, error)
}
  
func RepositoryTicket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTicket(ticket models.Ticket)(models.Ticket, error){
	err := r.db.Preload("station").Create(&ticket).Error

	return ticket, err
}

func (r *repository) FindTicket()([]models.Ticket, error){
	var ticket []models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").Find(&ticket).Error

	return ticket, err
}

func (r *repository) GetTicket(ID int)(models.Ticket, error){
	var ticket models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").First(&ticket, ID).Error

	return ticket, err
}