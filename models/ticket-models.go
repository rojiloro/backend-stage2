package models

import dto ""

type Ticket struct {
	ID        int    `json:"id" gorm:"primarykey:autoIncrement"`
	NameTrain string `json:"name_train" gorm:"type: varchar(255)"`
	TypeTrain string `json:"type_train" gorm:"type: varchar(255)"`
	StartDate string `json:"Start_date" gorm:"type: Date"`
	StartStationID StationResponse	`json:"start_station_id" gorm: "type: "`
}