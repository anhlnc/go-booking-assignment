package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID          uuid.UUID
	Customer_id uuid.UUID
	Flight_id   uuid.UUID
	Code        string
	Status      string
	Booked_date string
}
