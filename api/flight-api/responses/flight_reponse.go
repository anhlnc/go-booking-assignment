package responses

type FlightResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	From           string
	To             string
	Date           string
	Status         string
	Available_slot int
}
