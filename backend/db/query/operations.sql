-- name: GetOperationsFromDeviceID :many
SELECT id, device_id, name, rly_on
FROM operations
WHERE device_id = ?;

-- name: GetOperationFromID :one
SELECT id, device_id, name, rly_on
FROM operations
WHERE id = ?;

-- name: CreateOperation :execlastid
INSERT INTO operations (device_id, name, rly_on)
VALUES (?, ?, ?);