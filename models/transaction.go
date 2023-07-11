package models

type Transaction struct {
	ID       int    `json:"id" gorm:"primarykey:autoIncrement"`
	UserId   int    `json:"user_id"`
	User     User   `json:"user"`
	TicketId int    `json:"ticket_id"`
	Ticket   Ticket `json:"ticket"`
	Image    string `json:"image"`
	Status   string `json:"status"`
}