package models

// station models struct
type Station struct {
	ID int `json:"id" gorm:"primarykey:autoIncrement"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}