package requests

type CreateFlightRequest struct {
	Name           string
	From           string
	To             string
	Date           string
	Status         string
	Available_slot int
}
type UpdateFlightRequest struct {
	Id             string
	Name           string
	From           string
	To             string
	Date           string
	Status         string
	Available_slot int
}

type FindFlightRequest struct {
	Name string
	From string
	To   string
	Date string
}
