package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	ID             uuid.UUID
	Name           string
	From           string
	To             string
	Date           string
	Status         string
	Available_slot int
}
