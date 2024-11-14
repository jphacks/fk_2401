-- name: GetNodesFromWorkflow :many
SELECT 
    id, workflow_id, type, data, position_x, position_y
FROM nodes
WHERE workflow_id = ?;

-- name: CreateNode :execlastid
INSERT INTO nodes (workflow_id, type, data, position_x, position_y) 
VALUES (?, ?, ?, ?, ?);
