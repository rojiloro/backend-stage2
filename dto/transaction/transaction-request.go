package transaction

type CreateTransactionRequest struct {
	UserId   int `json:"user_id"`
	TicketId int `json:"ticket_id"`
}