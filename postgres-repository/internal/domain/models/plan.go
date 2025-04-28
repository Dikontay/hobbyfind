package models

import "time"

type Plan struct {
	StandardProperties
	Type     string  `json:"type" bun:"type,unique,notnull"`
	Price    float32 `json:"price" bun:"price,notnull"`        // e.g., 19.99
	Currency string  `json:"currency" bun:"currency,notnull"`  // e.g., "USD"
	StripeID string  `json:"stripe_id" bun:"stripe_id,unique"` // Link to Stripe Price ID
}

func (p Plan) OrderColumn() string {
	return "created_at"
}

func (p Plan) SetId(id string) {
	p.ID = id
}

func (p Plan) GetId() string {
	return p.ID
}

func (p Plan) GetSortConditions() map[string][]interface{} {
	conditions := make(map[string][]interface{})
	if p.Type != "" {
		conditions["type = ?"] = []interface{}{p.Type}
	}
	if p.Price != 0 {
		conditions["price = ?"] = []interface{}{p.Price}
	}
	if p.Currency != "" {
		conditions["currency = ?"] = []interface{}{p.Currency}
	}
	if p.StripeID != "" {
		conditions["stripe_id = ?"] = []interface{}{p.StripeID}
	}

	return conditions
}

func (p Plan) GetUpdateColumns() []string {
	columns := make([]string, 0)

	now := time.Now()
	p.UpdatedAt = &now
	columns = append(columns, "updated_at")

	if p.Type != "" {
		columns = append(columns, "type")
	}
	if p.Price != 0 {
		columns = append(columns, "price")
	}
	if p.Currency != "" {
		columns = append(columns, "currency")
	}
	if p.StripeID != "" {
		columns = append(columns, "stripe_id")
	}

	return columns
}
