package models

import "time"

type Instructor struct {
	StandardProperties
	AccountID      string  `json:"account_id" bun:"account_id,unique,notnull"`
	OrganizationID *string `json:"organization_id,omitempty" bun:"organization_id,type:uuid,nullzero"` // Nullable
	Bio            string  `json:"bio" bun:"bio"`
	Expertise      string  `json:"expertise" bun:"expertise"`
	Rating         float32 `json:"rating" bun:"rating,default:0"`
}

func (i Instructor) OrderColumn() string {
	return "created_at"
}

func (i Instructor) SetId(id string) {
	i.ID = id
}

func (i Instructor) GetId() string {
	return i.ID
}

func (i Instructor) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})
	if i.AccountID != "" {
		conditions["account_id = ?"] = []interface{}{i.AccountID}
	}
	if i.OrganizationID != nil {
		conditions["organization_id = ?"] = []interface{}{*i.OrganizationID}
	}
	if i.Bio != "" {
		conditions["bio = ?"] = []interface{}{i.Bio}
	}
	if i.Expertise != "" {
		conditions["expertise = ?"] = []interface{}{i.Expertise}
	}
	if i.Rating != 0 {
		conditions["rating = ?"] = []interface{}{i.Rating}
	}

	return conditions
}

func (i Instructor) GetUpdateColumns() []string {
	columns := make([]string, 0)
	now := time.Now()
	i.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if i.AccountID != "" {
		columns = append(columns, "account_id")
	}
	if i.OrganizationID != nil {
		columns = append(columns, "organization_id")
	}

	if i.Bio != "" {
		columns = append(columns, "bio")
	}
	if i.Expertise != "" {
		columns = append(columns, "expertise")
	}

	if i.Rating != 0 {
		columns = append(columns, "rating")
	}
	return columns
}
