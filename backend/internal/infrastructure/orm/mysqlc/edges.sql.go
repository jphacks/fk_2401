// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: edges.sql

package mysqlc

import (
	"context"
)

const createEdge = `-- name: CreateEdge :execlastid
INSERT INTO edges (workflow_id, source_node_id, target_node_id) 
VALUES (?, ?, ?)
`

type CreateEdgeParams struct {
	WorkflowID   int32
	SourceNodeID string
	TargetNodeID string
}

func (q *Queries) CreateEdge(ctx context.Context, arg CreateEdgeParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createEdge, arg.WorkflowID, arg.SourceNodeID, arg.TargetNodeID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getEdgesFromWorkflow = `-- name: GetEdgesFromWorkflow :many
SELECT 
    id, workflow_id, source_node_id, target_node_id
FROM edges
WHERE workflow_id = ?
`

type GetEdgesFromWorkflowRow struct {
	ID           int32
	WorkflowID   int32
	SourceNodeID string
	TargetNodeID string
}

func (q *Queries) GetEdgesFromWorkflow(ctx context.Context, workflowID int32) ([]GetEdgesFromWorkflowRow, error) {
	rows, err := q.db.QueryContext(ctx, getEdgesFromWorkflow, workflowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEdgesFromWorkflowRow
	for rows.Next() {
		var i GetEdgesFromWorkflowRow
		if err := rows.Scan(
			&i.ID,
			&i.WorkflowID,
			&i.SourceNodeID,
			&i.TargetNodeID,
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
