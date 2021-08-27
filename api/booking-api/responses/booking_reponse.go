package responses

import "time"

type BookingResponse struct {
	Code         string `json:"code"`
	Booking_Date string
	Customer_id  string
	Flight_id    string
	Status       string
}

type DeleteBookingResponse struct {
	Code        string
	Cancel_date time.Time
	Customer_id string
	Flight_id   string
	Status      string
}
