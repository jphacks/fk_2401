// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: houses.sql

package mysqlc

import (
	"context"
)

const createHouse = `-- name: CreateHouse :execlastid
INSERT INTO houses (name) 
VALUES (?)
`

func (q *Queries) CreateHouse(ctx context.Context, name string) (int64, error) {
	result, err := q.db.ExecContext(ctx, createHouse, name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getAllHouse = `-- name: GetAllHouse :many
SELECT id, name, created_at, updated_at FROM houses
`

func (q *Queries) GetAllHouse(ctx context.Context) ([]House, error) {
	rows, err := q.db.QueryContext(ctx, getAllHouse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []House
	for rows.Next() {
		var i House
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
