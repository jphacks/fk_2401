-- name: GetDeviceConditionsFromDeviceID :many
SELECT id, device_id, operation_id, valid, set_point, duration, operator
FROM device_conditions
WHERE device_id = ?;

-- name: GetDeviceConditionFromID :one
SELECT id, device_id, operation_id, valid, set_point, duration, operator
FROM device_conditions
WHERE id = ?;

-- name: CreateDeviceCondition :execlastid
INSERT INTO device_conditions (device_id, operation_id, valid, set_point, duration, operator)
VALUES (?, ?, ?, ?, ?, ?);