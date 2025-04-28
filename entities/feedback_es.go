package models

import "time"

type FeedbackES struct {
	ID               string    `json:"id"`
	OrganizationID   string    `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	UserName         string    `json:"user_name"`
	Text             string    `json:"text"`
	Rating           float64   `json:"rating"`
	CreatedAt        time.Time `json:"created_at"`
}
