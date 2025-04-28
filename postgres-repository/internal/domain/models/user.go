package models

import (
	"time"
)

type User struct {
	StandardProperties
	AccountId     string    `json:"account_id" bun:"account_id,type:uuid"`
	FullName      string    `json:"full_name" bun:"full_name"`
	Email         string    `json:"email" bun:"email,unique,notnull"`
	Phone         string    `json:"phone" bun:"phone"`
	Username      string    `json:"username" bun:"username,unique,notnull"`
	Password      string    `json:"password" bun:"password,notnull"`
	PlanID        string    `json:"plan_id" bun:"plan_id,type:uuid"`
	PlanStartDate time.Time `json:"plan_start_date" bun:"plan_start_date,nullzero"`
	PlanEndDate   time.Time `json:"plan_end_date" bun:"plan_end_date,nullzero"`
	IsActive      bool      `json:"is_active" bun:"is_active,default:true"`
}

func (u User) OrderColumn() string {
	return "created_at"
}

func (u User) SetId(id string) {
	u.ID = id
}

func (u User) GetId() string {
	return u.ID
}

func (u User) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})
	if u.Phone != "" {
		conditions["phone = ?"] = []interface{}{u.Phone}
	}
	if u.FullName != "" {
		conditions["full_name = ?"] = []interface{}{u.FullName}
	}
	if u.Email != "" {
		conditions["email = ?"] = []interface{}{u.Email}
	}
	if u.Username != "" {
		conditions["username = ?"] = []interface{}{u.Username}
	}
	if u.Password != "" {
		conditions["password = ?"] = []interface{}{u.Password}
	}
	if u.AccountId != "" {
		conditions["account_id = ?"] = []interface{}{u.AccountId}
	}
	if u.PlanID != "" {
		conditions["plan_id = ?"] = []interface{}{u.PlanID}
	}
	if u.PlanStartDate != (time.Time{}) {
		conditions["plan_start_date = ?"] = []interface{}{u.PlanStartDate}
	}
	if u.PlanEndDate != (time.Time{}) {
		conditions["plan_end_date = ?"] = []interface{}{u.PlanEndDate}
	}
	if u.IsActive {
		conditions["is_active = ?"] = []interface{}{u.IsActive}
	}

	return conditions
}

func (u User) GetUpdateColumns() []string {

	columns := make([]string, 0)

	now := time.Now()
	u.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if u.AccountId != "" {
		columns = append(columns, "account_id")
	}
	if u.FullName != "" {
		columns = append(columns, "full_name")
	}
	if u.Email != "" {
		columns = append(columns, "email")
	}
	if u.Phone != "" {
		columns = append(columns, "phone")
	}
	if u.Username != "" {
		columns = append(columns, "username")
	}
	if u.Password != "" {
		columns = append(columns, "password")
	}

	if u.PlanID != "" {
		columns = append(columns, "plan_id")
	}
	if u.PlanStartDate != (time.Time{}) {
		columns = append(columns, "plan_start_date")
	}
	if u.PlanEndDate != (time.Time{}) {
		columns = append(columns, "plan_end_date")
	}
	if u.IsActive {
		columns = append(columns, "is_active")
	}
	return columns
}
