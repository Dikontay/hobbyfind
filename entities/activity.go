package models

import "time"

type Activity struct {
	StandardProperties
	Name        string `json:"name" bun:"name,unique,notnull"`
	Description string `json:"description" bun:"description"`
}

func (a Activity) OrderColumn() string {
	return "created_at"
}

func (a Activity) SetId(id string) {
	a.ID = id
}

func (a Activity) GetId() string {
	return a.ID
}

func (a Activity) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})
	if a.Name != "" {
		conditions["name = ?"] = []interface{}{a.Name}
	}
	if a.Description != "" {
		conditions["description = ?"] = []interface{}{a.Description}
	}

	return conditions
}

func (a Activity) GetUpdateColumns() []string {
	columns := make([]string, 0)

	now := time.Now()
	a.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if a.Name != "" {
		columns = append(columns, "name")
	}
	if a.Description != "" {
		columns = append(columns, "description")
	}

	return columns
}
