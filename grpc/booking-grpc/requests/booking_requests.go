package requests

import "github.com/gofrs/uuid"

type UpdateBookingRequest struct {
	ID          uuid.UUID
	Customer_id string
	Flight_id   string
	Code        string
	Status      string
	Booked_date string
}

type ViewBookingRequest struct {
	Customer_id string
	Code        string
}

type ListBookingRequest struct {
	Customer_id string
}
