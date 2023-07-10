package transaction

type TransactionResponse struct {
	ID       int `json:"id"`
	UserId   int `json:"user_id"`
	TicketId int `json:"ticket_id"`
}