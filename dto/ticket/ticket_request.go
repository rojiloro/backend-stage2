package ticketdto

type CreateTicketRequest struct {
	NameTrain            string `json:"name_train" form:"name_train" validate:"required"`
	TypeTrain            string `json:"type_train" form:"type_train" validate:"required"`
	StartDate            string `json:"start_date_train" form:"start_date_train" validate:"required"`
	StartStationID       int    `json:"start_station_id" form:"start_station_id" validate:"required"`
	StartTime            string `json:"start_time" form:"start_time" validate:"required"`
	DestinationStationID int    `json:"destination_station_id" form:"destination_station_id" validate:"required"`
	ArrivalTime          string `json:"arrival_time" form:"arrival_time" validate:"required"`
	Price                int    `json:"price" form:"price" validate:"required"`
	Qty                  int    `json:"qty" form:"qty" validate:"required"`
}

type UpdateTicketRequest struct {
	NameTrain            string `json:"name_train" form:"name_train"`
	TypeTrain            string `json:"type_train" form:"type_train"`
	StartDate            string `json:"start_date_train" form:"start_date_train"`
	StartStationID       int    `json:"start_station_id" form:"start_station_id"`
	StartTime            string `json:"start_time" form:"start_time"`
	DestinationStationID int    `json:"destination_station_id" form:"destination_station_id" validate:"required"`
	ArrivalTime          string `json:"arrival_time" form:"arrival_time"`
	Price                int    `json:"price" form:"price"`
	Qty                  int    `json:"qty" form:"qty"`
}