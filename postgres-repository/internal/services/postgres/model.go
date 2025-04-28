package postgres

import "time"

type Model interface {
	OrderColumn() string
	SetId(id string)
	GetId() string
	GetSortConditions() map[string][]interface{}
	GetUpdateColumns() []string
}

type DateSortProperties struct {
	CreatedBefore *time.Time `bun:"-" json:"created_before,omitempty" example:"2020-01-01T00:00:00Z"`
	CreatedAfter  *time.Time `bun:"-" json:"created_after,omitempty"  example:"2020-01-01T00:00:00Z"`

	UpdatedBefore *time.Time `bun:"-" json:"updated_before,omitempty" example:"2020-01-01T00:00:00Z"`
	UpdatedAfter  *time.Time `bun:"-" json:"updated_after,omitempty" example:"2020-01-01T00:00:00Z"`

	DeletedBefore *time.Time `bun:"-" json:"deleted_before,omitempty" example:"2020-01-01T00:00:00Z"`
	DeletedAfter  *time.Time `bun:"-" json:"deleted_after,omitempty" example:"2020-01-01T00:00:00Z"`
}

func (r DateSortProperties) GetSortConditions(conditions map[string][]interface{}) map[string][]interface{} {
	if r.CreatedBefore != nil {
		conditions["created_at <= ?"] = []interface{}{r.CreatedBefore}
	}
	if r.CreatedAfter != nil {
		conditions["created_at >= ?"] = []interface{}{r.CreatedAfter}
	}
	if r.UpdatedBefore != nil {
		conditions["updated_at <= ?"] = []interface{}{r.UpdatedBefore}
	}
	if r.UpdatedAfter != nil {
		conditions["updated_at >= ?"] = []interface{}{r.UpdatedAfter}
	}
	if r.DeletedBefore != nil {
		conditions["deleted_at <= ?"] = []interface{}{r.DeletedBefore}
	}
	if r.DeletedAfter != nil {
		conditions["deleted_at >= ?"] = []interface{}{r.DeletedAfter}
	}

	return conditions
}

type StandardProperties struct {
	CreatedAt *time.Time `json:"created_at" bun:"created_at,default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bun:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`
}

func (r *StandardProperties) SetUpdatedAt(time time.Time) {
	r.UpdatedAt = &time
}
