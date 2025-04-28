package models

import "time"

type Class struct {
	StandardProperties
	ActivityID        string    `json:"activity_id" bun:"activity_id,type:uuid,notnull"`
	OrganizationID    string    `json:"organization_id" bun:"organization_id,type:uuid,notnull"`
	StartTime         time.Time `json:"start_time" bun:"start_time,notnull"`
	EndTime           time.Time `json:"end_time" bun:"end_time,notnull"`
	IsRecurring       bool      `json:"is_recurring" bun:"is_recurring,default:false"`
	RecurrencePattern string    `json:"recurrence_pattern" bun:"recurrence_pattern,nullzero"`
	TotalNumber       int       `json:"total_number" bun:"total_number,notnull"`
	CurrentNumber     int       `json:"current_number" bun:"current_number,default:0"`
}

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
