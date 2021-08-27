package requests

import "github.com/gofrs/uuid"

type UpdateCustomertRequest struct {
	ID           uuid.UUID
	Name         string
	Address      string
	License_id   string
	Phone_number string
	Email        string
	Password     string
	Active       bool
}

type ListCustomerRequest struct {
	Name string
}
