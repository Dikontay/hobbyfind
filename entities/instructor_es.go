package models

type InstructorES struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Bio              string   `json:"bio"`
	Expertise        []string `json:"expertise"`         // e.g., ["Piano", "Music Theory"]
	OrganizationName string   `json:"organization_name"` // Nullable if freelancer
	City             string   `json:"city"`
	Rating           float64  `json:"rating"`
}
