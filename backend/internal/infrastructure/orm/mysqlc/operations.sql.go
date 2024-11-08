// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: operations.sql

package mysqlc

import (
	"context"
)

const createOperation = `-- name: CreateOperation :execlastid
INSERT INTO operations (device_id, name, rly_on)
VALUES (?, ?, ?)
`

type CreateOperationParams struct {
	DeviceID int32
	Name     string
	RlyOn    int32
}

func (q *Queries) CreateOperation(ctx context.Context, arg CreateOperationParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createOperation, arg.DeviceID, arg.Name, arg.RlyOn)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getOperationFromID = `-- name: GetOperationFromID :one
SELECT id, device_id, name, rly_on
FROM operations
WHERE id = ?
`

type GetOperationFromIDRow struct {
	ID       int32
	DeviceID int32
	Name     string
	RlyOn    int32
}

func (q *Queries) GetOperationFromID(ctx context.Context, id int32) (GetOperationFromIDRow, error) {
	row := q.db.QueryRowContext(ctx, getOperationFromID, id)
	var i GetOperationFromIDRow
	err := row.Scan(
		&i.ID,
		&i.DeviceID,
		&i.Name,
		&i.RlyOn,
	)
	return i, err
}

const getOperationsFromDeviceID = `-- name: GetOperationsFromDeviceID :many
SELECT id, device_id, name, rly_on
FROM operations
WHERE device_id = ?
`

type GetOperationsFromDeviceIDRow struct {
	ID       int32
	DeviceID int32
	Name     string
	RlyOn    int32
}

func (q *Queries) GetOperationsFromDeviceID(ctx context.Context, deviceID int32) ([]GetOperationsFromDeviceIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getOperationsFromDeviceID, deviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOperationsFromDeviceIDRow
	for rows.Next() {
		var i GetOperationsFromDeviceIDRow
		if err := rows.Scan(
			&i.ID,
			&i.DeviceID,
			&i.Name,
			&i.RlyOn,
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
