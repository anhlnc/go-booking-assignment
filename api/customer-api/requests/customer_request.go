package requests

type CreateCustomerRequest struct {
	Name         string
	Address      string
	License_id   string
	Phone_number string
	Email        string
	Password     string
	Active       bool
}
type UpdateCustomerRequest struct {
	Id           string
	Name         string
	Address      string
	License_id   string
	Phone_number string
	Email        string
	Password     string
	Active       bool
}
type BookingHistoryRequest struct {
	Id string
}
type ChangePasswordRequest struct {
	Id       string
	Password string
}
