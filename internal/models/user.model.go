package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	FirstName    string    `gorm:"type:varchar(255)" json:"first_name"`
	LastName     string    `gorm:"type:varchar(255)" json:"last_name"`
	Avatar       string    `gorm:"type:text" json:"avatar"`
	PasswordHash string    `gorm:"type:text" json:"password"`
	gorm.Model
}
