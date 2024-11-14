-- name: GetNodesFromWorkflow :many
SELECT 
    id, workflows_id, type, data, position_x, position_y
FROM nodes
WHERE workflows_id = ?;
