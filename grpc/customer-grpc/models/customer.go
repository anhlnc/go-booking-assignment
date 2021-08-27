package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID           uuid.UUID
	Name         string
	Address      string
	License_id   string
	Phone_number string
	Email        string
	Password     string
	Active       bool
}
