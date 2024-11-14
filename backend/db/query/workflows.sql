-- name: GetAllWorkflows :many
SELECT id, name
FROM workflows;

-- name: CreateWorkflow :execlastid
INSERT INTO workflows (name) 
VALUES (?);
