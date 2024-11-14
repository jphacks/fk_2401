-- name: GetEdgesFromWorkflow :many
SELECT 
    id, workflow_id, source_node_id, target_node_id
FROM edges
WHERE workflow_id = ?;

-- name: CreateEdge :execlastid
INSERT INTO edges (workflow_id, source_node_id, target_node_id) 
VALUES (?, ?, ?);
