package ticketdto

type TicketResponse struct {
	ID                   int    `json:"id"`
	NameTrain            string `json:"name_train"`
	TypeTrain            string `json:"type_train"`
	StartDate            string `json:"start_date_train"`
	StartStationID       int    `json:"start_station_id"`
	StartTime            string `json:"start_time" `
	DestinationStationID int    `json:"destination_station_id"`
	ArrivalTime          string `json:"arrival_time"`
	Price                int    `json:"price"`
	Qty                  int    `json:"qty"`
	UserID               int    `json:"user_id"`
}
