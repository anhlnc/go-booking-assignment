package requests

type AssignBookingRequest struct {
	Customer_id string
	Flight_id   string
	Code        string
	Status      string
}
type ViewBookingRequest struct {
	Customer_id string
	Code        string
}
type DeleteBookingRequest struct {
	Id string
}
