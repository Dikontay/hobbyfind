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

func (c Class) OrderColumn() string {
	return "start_time"
}

func (c Class) SetId(id string) {
	c.ID = id
}

func (c Class) GetId() string {
	return c.ID
}

func (c Class) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})
	if c.ActivityID != "" {
		conditions["activity_id = ?"] = []interface{}{c.ActivityID}
	}
	if c.OrganizationID != "" {
		conditions["organization_id = ?"] = []interface{}{c.OrganizationID}
	}
	if !c.StartTime.IsZero() {
		conditions["start_time = ?"] = []interface{}{c.StartTime}
	}
	if !c.EndTime.IsZero() {
		conditions["end_time = ?"] = []interface{}{c.EndTime}
	}
	if c.IsRecurring {
		conditions["is_recurring = ?"] = []interface{}{c.IsRecurring}
	}
	if c.RecurrencePattern != "" {
		conditions["recurrence_pattern = ?"] = []interface{}{c.RecurrencePattern}
	}
	if c.TotalNumber > 0 {
		conditions["total_number = ?"] = []interface{}{c.TotalNumber}
	}
	if c.CurrentNumber > 0 {
		conditions["current_number = ?"] = []interface{}{c.CurrentNumber}
	}

	return conditions
}

func (c Class) GetUpdateColumns() []string {
	columns := make([]string, 0)

	now := time.Now()
	c.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if c.ActivityID != "" {
		columns = append(columns, "activity_id")
	}
	if c.OrganizationID != "" {
		columns = append(columns, "organization_id")
	}
	if !c.StartTime.IsZero() {
		columns = append(columns, "start_time")
	}
	if !c.EndTime.IsZero() {
		columns = append(columns, "end_time")
	}
	if c.IsRecurring {
		columns = append(columns, "is_recurring")
	}
	if c.RecurrencePattern != "" {
		columns = append(columns, "recurrence_pattern")
	}
	if c.TotalNumber > 0 {
		columns = append(columns, "total_number")
	}
	if c.CurrentNumber > 0 {
		columns = append(columns, "current_number")
	}

	return columns
}
