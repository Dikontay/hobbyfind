package models

import "time"

type ClassES struct {
	ID               string    `json:"id"`    // Use Postgres UUID
	Title            string    `json:"title"` // If you have class titles
	Description      string    `json:"description"`
	Activity         string    `json:"activity"` // e.g., "Yoga"
	City             string    `json:"city"`
	OrganizationName string    `json:"organization_name"`
	InstructorName   string    `json:"instructor_name"`
	StartTime        time.Time `json:"start_time"`
	IsRecurring      bool      `json:"is_recurring"`
	Rating           float64   `json:"rating"` // Precomputed avg rating
}
