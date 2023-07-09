package models

type Ticket struct {
	ID                   int     `json:"id" gorm:"primarykey:autoIncrement"`
	NameTrain            string  `json:"name_train" gorm:"type: varchar(255)"`
	TypeTrain            string  `json:"type_train" gorm:"type: varchar(255)"`
	StartDate            string  `json:"start_date" gorm:"type: varchar(255)"`
	StartStationID       int     `json:"start_station_id"`
	StartStation         Station `json:"start_station"`
	StartTime            string  `json:"start_time" gorm:"type: varchar(255)"`
	DestinationStationID int     `json:"destination_station_id"`
	DestinationStation   Station `json:"destination_station"`
	ArrivalTime          string  `json:"arrival_time" gorm:"type: varchar(255)"`
	Price                int     `json:"price" gorm:"type : int"`
	Qty                  int     `json:"qty" gorm:"type : int"`
}