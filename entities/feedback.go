package models

import "time"

type Feedback struct {
	StandardProperties
	UserID         string    `json:"user_id" bun:"user_id,type:uuid,notnull"`
	OrganizationID string    `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	Text           string    `json:"text" bun:"text,notnull"`
	Rating         float32   `json:"rating" bun:"rating,notnull"` // e.g., 1.0 to 5.0
	CreatedAt      time.Time `json:"created_at" bun:"created_at,default:now()"`
}
