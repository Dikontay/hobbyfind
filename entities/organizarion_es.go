package models

type OrganizationES struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	Activities  []string `json:"activities"` // e.g., ["Music", "Dance"]
	Rating      float64  `json:"rating"`
}
