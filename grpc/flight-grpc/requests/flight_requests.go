package requests

import "github.com/gofrs/uuid"

type UpdateFlightRequest struct {
	Id             uuid.UUID
	Name           string
	From           string
	To             string
	Date           string
	Status         string
	Available_slot int
}

type ListFlightRequest struct {
	Name string
}

type FindFlightRequest struct {
	Name string
	From string
	To   string
	Date string
}
