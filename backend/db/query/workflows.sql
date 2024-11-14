-- name: GetAllWorkflows :many
SELECT id, name, time_schedules_id
FROM workflows;
