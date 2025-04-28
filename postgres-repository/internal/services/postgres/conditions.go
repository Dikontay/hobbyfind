package postgres

import (
	"fmt"
	"github.com/uptrace/bun"
	"strings"
)

func ConditionId(property, value string, conditions map[string][]interface{}) {
	value = strings.ReplaceAll(value, " ", "")
	if len(value) == 0 {
		return
	}
	if value == "null" {
		conditions[fmt.Sprintf("%s IS NULL", property)] = nil
		return
	}

	fields := strings.Split(value, ",")
	if fieldsCount := len(fields); fieldsCount > 1 {
		conditions[fmt.Sprintf("%s IN (?)", property)] = []interface{}{bun.In(fields)}
	} else if fieldsCount == 1 {
		conditions[fmt.Sprintf("%s = ?", property)] = []interface{}{value}
	}
}

func ConditionStrLike(property, value string, conditions map[string][]interface{}) {
	value = strings.TrimSpace(value)
	property = strings.TrimSpace(property)

	if len(value) > 0 {
		conditions[fmt.Sprintf("%s like ?", property)] = []interface{}{fmt.Sprintf("%%%s%%", value)}
	}
}

func ConditionStrEqual(property string, value string, conditions map[string][]interface{}) {
	property = strings.TrimSpace(property)

	if len(value) > 0 {
		conditions[fmt.Sprintf("%s = ?", property)] = []interface{}{value}
	}
}
