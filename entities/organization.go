package models

import (
	"time"
)

type Organization struct {
	StandardProperties
	Name        string  `json:"name" bun:"name,notnull"`
	Description string  `json:"description" bun:"description"`
	Email       string  `json:"email" bun:"email,unique,notnull"`
	Phone       string  `json:"phone" bun:"phone,notnull"`
	Username    string  `json:"username" bun:"username,unique,notnull"`
	Password    string  `json:"password" bun:"password,notnull"`
	LocationID  string  `json:"location_id" bun:"location_id,type:uuid,notnull"`
	Rating      float32 `json:"rating" bun:"rating,default:0"`
}

func (o Organization) OrderColumn() string {
	return "created_at"
}

func (o Organization) SetId(id string) {
	o.ID = id
}

func (o Organization) GetId() string {
	return o.ID
}

func (o Organization) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})
	if o.Name != "" {
		conditions["name = ?"] = []interface{}{o.Name}
	}
	if o.Description != "" {
		conditions["description = ?"] = []interface{}{o.Description}
	}
	if o.Email != "" {
		conditions["email = ?"] = []interface{}{o.Email}
	}
	if o.Phone != "" {
		conditions["phone = ?"] = []interface{}{o.Phone}
	}
	if o.Username != "" {
		conditions["username = ?"] = []interface{}{o.Username}
	}
	if o.Password != "" {
		conditions["password = ?"] = []interface{}{o.Password}
	}
	if o.LocationID != "" {
		conditions["location_id = ?"] = []interface{}{o.LocationID}
	}
	if o.Rating != 0 {
		conditions["rating = ?"] = []interface{}{o.Rating}
	}

	return conditions
}

func (o Organization) GetUpdateColumns() []string {
	columns := make([]string, 0)

	now := time.Now()
	o.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if o.Name != "" {
		columns = append(columns, "name")
	}
	if o.Description != "" {
		columns = append(columns, "description")
	}
	if o.Email != "" {
		columns = append(columns, "email")
	}
	if o.Phone != "" {
		columns = append(columns, "phone")
	}
	if o.Username != "" {
		columns = append(columns, "username")
	}
	if o.Password != "" {
		columns = append(columns, "password")
	}
	if o.LocationID != "" {
		columns = append(columns, "location_id")
	}
	if o.Rating != 0 {
		columns = append(columns, "rating")
	}

	return columns
}
