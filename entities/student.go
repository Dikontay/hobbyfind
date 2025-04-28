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
