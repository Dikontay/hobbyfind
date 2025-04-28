package models

import "time"

type StandardProperties struct {
	ID        string     `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt *time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bun:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bun:"deleted_at"`
}
