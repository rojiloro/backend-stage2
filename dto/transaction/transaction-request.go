package transaction

type CreateTransactionRequest struct {
	UserId   int    `json:"user_id" form:"user_id" validate:"required"`
	TicketId int    `json:"ticket_id" form:"ticket_id" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
}