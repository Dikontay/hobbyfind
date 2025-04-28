package models

type Instructor struct {
	StandardProperties
	AccountID      string  `json:"account_id" bun:"account_id,unique,notnull"`
	OrganizationID *string `json:"organization_id,omitempty" bun:"organization_id,type:uuid,nullzero"` // Nullable
	Bio            string  `json:"bio" bun:"bio"`
	Expertise      string  `json:"expertise" bun:"expertise"`
	Rating         float32 `json:"rating" bun:"rating,default:0"`
}

type InstructorES struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Bio              string   `json:"bio"`
	Expertise        []string `json:"expertise"`         // e.g., ["Piano", "Music Theory"]
	OrganizationName string   `json:"organization_name"` // Nullable if freelancer
	City             string   `json:"city"`
	Rating           float64  `json:"rating"`
}
