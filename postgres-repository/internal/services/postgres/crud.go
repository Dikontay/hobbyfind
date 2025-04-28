package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"reflect"
)

func ExistsByPrimaryKey[M Model](obj M) (bool, error) {
	if cn == nil {
		return false, fmt.Errorf("connection service is empty, use Init() function")
	}

	ctx := context.Background()

	db, err := cn.GetClient()
	if err != nil {
		return false, fmt.Errorf("failed to get connection: %v", err)
	}

	exists, err := db.NewSelect().Model(obj).WherePK().Exists(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %v", err)
	}

	return exists, nil
}

func ExistsByColumn[M Model](obj M) (bool, error) {
	if cn == nil {
		return false, fmt.Errorf("connection service is empty, use Init() function")
	}

	ctx := context.Background()

	db, err := cn.GetClient()
	if err != nil {
		return false, fmt.Errorf("failed to get connection: %v", err)
	}

	query := db.NewSelect().Model(obj)

	exists, err := query.WhereOr(query.String(), obj.GetSortConditions()).Exists(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %v", err)
	}

	return exists, nil

}

func Create[M Model](obj M) error {
	if cn == nil {
		return fmt.Errorf("connection service is empty, use Init() function")
	}
	ctx := context.Background()

	db, err := cn.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get connection: %v", err)
	}

	_, err = db.NewInsert().Model(obj).Returning("id, created_at").Exec(ctx)
	if err != nil {
		return fmt.Errorf("execute query level: %v", err)
	}

	return nil
}
func Find[M Model](obj M, relations ...string) error {
	if cn == nil {
		return fmt.Errorf("connection service is empty, use Init() function")
	}
	value := reflect.ValueOf(obj)

	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	db, err := cn.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get connection: %v", err)
	}

	ctx := context.Background()
	query := db.NewSelect().Model(obj).WherePK()

	for _, relation := range relations {
		query = query.Relation(relation)
	}

	err = query.Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return fmt.Errorf("execute query error: %v", err)
	}

	return nil
}
func List[M Model](obj M, pageNum, pageCount int, relations ...string) ([]M, error) {
	relationMap := make(map[string]map[string][]interface{})
	for _, relation := range relations {
		relationMap[relation] = make(map[string][]interface{})
	}
	objList, err := ListQueryExec[M](pageNum, pageCount, obj.OrderColumn(), obj.GetSortConditions(), relationMap)
	if err != nil {
		return nil, err
	}
	return objList, nil
}

func FullList[M Model](obj M, relations ...string) ([]M, error) {
	if cn == nil {
		return nil, fmt.Errorf("connection service is empty, use Init() function")
	}
	db, err := cn.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %v", err)
	}

	usersList := make([]M, 0)

	query := db.NewSelect().Model(&usersList).Order(obj.OrderColumn())
	conditions := obj.GetSortConditions()

	for _, relation := range relations {
		query.Relation(relation)
	}

	for conditionQuery, values := range conditions {
		query = query.Where(conditionQuery, values...)
	}
	ctx := context.Background()

	err = query.Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("execute query error: %v", err)
	}

	return usersList, nil
}
func Count[M Model](obj M) (int, error) {
	if cn == nil {
		return 0, fmt.Errorf("connection service is empty, use Init() function")
	}
	db, err := cn.GetClient()
	if err != nil {
		return 0, fmt.Errorf("failed to get connection: %v", err)
	}

	query := db.NewSelect().Model(obj)

	conditions := obj.GetSortConditions()

	for conditionQuery, values := range conditions {
		query = query.Where(conditionQuery, values...)
	}

	ctx := context.Background()

	count, err := query.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("execute query error: %v", err)
	}

	return count, nil
}
func Update[M Model](obj M) error {
	if cn == nil {
		return fmt.Errorf("connection service is empty, use Init() function")
	}
	db, err := cn.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get connection: %v", err)
	}

	value := reflect.ValueOf(obj)

	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	ctx := context.Background()

	columns := obj.GetUpdateColumns()

	_, err = db.NewUpdate().Model(obj).Returning("*").Column(columns...).WherePK().Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute query (%s): %v", obj.GetId(), err)
	}

	return nil
}
func Delete[M Model](obj M, permanent ...bool) error {
	if cn == nil {
		return fmt.Errorf("connection service is empty, use Init() function")
	}

	db, err := cn.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get connection: %v", err)
	}

	value := reflect.ValueOf(obj)

	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	ctx := context.Background()

	query := db.NewDelete().Model(obj).Returning("*").WherePK()

	if len(permanent) > 0 && permanent[0] == true {
		query = query.ForceDelete()
	}

	_, err = query.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute query (%s): %v", obj.GetId(), err)
	}

	return nil
}

func BulkCreate[M Model](objs []M) ([]M, error) {
	if cn == nil {
		return nil, fmt.Errorf("connection service is empty, use Init() function")
	}

	db, err := cn.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %v", err)
	}

	ctx := context.Background()

	_, err = db.NewInsert().Model(&objs).Returning("id, created_at").Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return objs, nil
}

func BulkUpdate[M Model](objs []M) ([]M, error) {
	if cn == nil {
		return nil, fmt.Errorf("connection service is empty, use Init() function")
	}

	db, err := cn.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %v", err)
	}

	updatedObjs := make([]M, 0, len(objs))

	var value reflect.Value
	err = db.RunInTx(context.Background(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, datum := range objs {
			value = reflect.ValueOf(datum)

			if value.Kind() != reflect.Ptr {
				return fmt.Errorf("obj in slice must be a pointer")
			}

			_, err = tx.NewUpdate().Model(datum).Column(datum.GetUpdateColumns()...).Returning("*").WherePK().Exec(ctx)
			if err != nil {
				return err
			}
			updatedObjs = append(updatedObjs, datum)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return updatedObjs, nil
}
