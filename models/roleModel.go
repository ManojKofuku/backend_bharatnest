package models


type Role struct {
	ID          string `gorm:"unique;primaryKey"`
	Role        string `gorm:"unique;not null"`
	Description string
}


