-- name: GetEdgesFromWorkflow :many
SELECT 
    id, workflows_id, source_node_id, target_node_id
FROM edges
WHERE workflows_id = ?;
