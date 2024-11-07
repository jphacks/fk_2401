-- name: GetAllSensor :many
SELECT id, ccm_type, room, region, `order`, `priority`
FROM sensors;

-- name: GetSensorFromID :one
SELECT id, ccm_type, room, region, `order`, `priority`
FROM sensors
WHERE id = ?;

-- name: CreateSensor :execlastid
INSERT INTO sensors (ccm_type, room, region, `order`, `priority`)
VALUES (?, ?, ?, ?, ?);