package postgres

import (
	"context"
	"fmt"
	"github.com/oleiade/reflections"
	"github.com/uptrace/bun"
	"reflect"
)

func ListQueryExec[Object any](pageNum, pageCount int, orderColumn string, conditions map[string][]interface{}, relationConditions map[string]map[string][]interface{}) ([]Object, error) {
	if cn == nil {
		return nil, fmt.Errorf("connection service is empty, use Init() function")
	}
	db, err := cn.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %v", err)
	}

	objList := make([]Object, 0)

	query := db.NewSelect().Model(&objList).Order(orderColumn)

	for relation, relConditions := range relationConditions {
		if relConditions == nil {
			query = query.Relation(relation)
		} else {
			query = query.Relation(relation, func(query *bun.SelectQuery) *bun.SelectQuery {
				for conditionQuery, values := range relConditions {
					query = query.Where(conditionQuery, values...)
				}
				return query
			})
		}
	}

	for conditionQuery, values := range conditions {
		query = query.Where(conditionQuery, values...)
	}
	ctx := context.Background()

	if pageCount > 0 {
		query = query.Limit(pageCount)
	}
	if pageNum > 0 {
		query = query.Offset((pageNum - 1) * pageCount)
	}

	err = query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("execute query error: %v", err)
	}

	objList = removeElements(objList, func(object Object) bool {
		for relation, cond := range relationConditions {
			if len(cond) > 0 {
				value, err := reflections.GetField(object, relation)
				if err != nil {
					return false
				}
				if reflect.ValueOf(value).IsZero() {
					return true
				}
			}
		}

		return false
	})

	return objList, nil
}

func removeElements[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice)) // создаём результирующий срез с начальной вместимостью
	for _, item := range slice {
		if !predicate(item) { // добавляем только те элементы, которые не удовлетворяют условию
			result = append(result, item)
		}
	}
	return result // возвращаем новый срез без нежелательных элементов
}
