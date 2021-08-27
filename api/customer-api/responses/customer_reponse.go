package responses

type CustomerResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	License_id   string
	Phone_number string
	Email        string
	Password     string
	Active       bool
}
