package models

import "time"

type Account struct {
	StandardProperties
	Username     string    `json:"username" bun:"username,unique,notnull"`
	Email        string    `json:"email" bun:"email,unique,notnull"`
	PasswordHash string    `json:"-" bun:"password_hash,notnull"`
	RoleType     string    `json:"role_type" bun:"role_type,notnull"` // enum: 'user', 'organization', 'admin', 'instructor', 'organization_admin'
	IsActive     bool      `json:"is_active" bun:"is_active,default:true"`
	LastLogin    time.Time `json:"last_login" bun:"last_login,nullzero"`
}

func (a Account) OrderColumn() string {
	return "created_at"
}

func (a Account) SetId(id string) {
	a.ID = id
}

func (a Account) GetId() string {
	return a.ID
}

func (a Account) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})

	if a.Username != "" {
		conditions["username = ?"] = []interface{}{a.Username}
	}
	if a.Email != "" {
		conditions["email = ?"] = []interface{}{a.Email}
	}
	if a.PasswordHash != "" {
		conditions["password_hash = ?"] = []interface{}{a.PasswordHash}
	}
	if a.RoleType != "" {
		conditions["role_type = ?"] = []interface{}{a.RoleType}
	}
	if a.IsActive {
		conditions["is_active = ?"] = []interface{}{a.IsActive}
	}
	if !a.LastLogin.IsZero() {
		conditions["last_login = ?"] = []interface{}{a.LastLogin}
	}

	return conditions
}

func (a Account) GetUpdateColumns() []string {
	columns := make([]string, 0)

	now := time.Now()
	a.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if a.Username != "" {
		columns = append(columns, "username")
	}
	if a.Email != "" {
		columns = append(columns, "email")
	}
	if a.PasswordHash != "" {
		columns = append(columns, "password_hash")
	}
	if a.RoleType != "" {
		columns = append(columns, "role_type")
	}
	if a.IsActive {
		columns = append(columns, "is_active")
	}
	if !a.LastLogin.IsZero() {
		columns = append(columns, "last_login")
	}

	return columns
}
