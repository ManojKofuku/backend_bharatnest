package models

type User struct {
	ID             string    `gorm:"unique;primaryKey"`
	FullName       string    `gorm:"not null"`
	Email          string    `gorm:"unique;not null"`
	PhoneNumber    string    `gorm:"unique;not null"`
	HashedPassword string    `gorm:"not null"`
	DateOfBirth    string
	IsInvited      bool      `gorm:"default:false"`
	RoleID         string
	Role           *Role     `gorm:"foreignKey:RoleID;references:ID"`
}




