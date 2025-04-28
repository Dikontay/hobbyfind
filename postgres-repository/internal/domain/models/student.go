package models

import "time"

type Student struct {
	StandardProperties
	ClassID    string    `json:"class_id" bun:"class_id,type:uuid,notnull"`
	UserID     string    `json:"user_id" bun:"user_id,type:uuid,notnull"`
	Status     string    `json:"status" bun:"status,default:'active'"`
	Attendance float32   `json:"attendance" bun:"attendance,default:0"`
	JoinedAt   time.Time `json:"joined_at" bun:"joined_at,default:now()"`
}

func (s Student) OrderColumn() string {
	return "joined_at"
}

func (s Student) SetId(id string) {
	s.ID = id
}

func (s Student) GetId() string {

	return s.ID
}

func (s Student) GetSortConditions() map[string][]interface{} {

	conditions := make(map[string][]interface{})
	if s.ClassID != "" {
		conditions["class_id = ?"] = []interface{}{s.ClassID}
	}
	if s.UserID != "" {
		conditions["user_id = ?"] = []interface{}{s.UserID}
	}
	if s.Status != "" {
		conditions["status = ?"] = []interface{}{s.Status}
	}
	if s.Attendance != 0 {
		conditions["attendance = ?"] = []interface{}{s.Attendance}
	}
	if !s.JoinedAt.IsZero() {
		conditions["joined_at = ?"] = []interface{}{s.JoinedAt}
	}

	return conditions
}

func (s Student) GetUpdateColumns() []string {
	columns := make([]string, 0)

	now := time.Now()
	s.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if s.ClassID != "" {
		columns = append(columns, "class_id")
	}
	if s.UserID != "" {
		columns = append(columns, "user_id")
	}
	if s.Status != "" {
		columns = append(columns, "status")
	}
	if s.Attendance != 0 {
		columns = append(columns, "attendance")
	}
	if !s.JoinedAt.IsZero() {
		columns = append(columns, "joined_at")
	}

	return columns
}
