-- name: GetAllUecsDevice :many
SELECT id, ccmtype, room, region, `order`, `priority`
FROM uecs_devices;

-- name: GetUecsDeviceFromID :one
SELECT id, ccmtype, room, region, `order`, `priority`
FROM uecs_devices
WHERE id = ?;

-- name: CreateUecsDevice :execlastid
INSERT INTO uecs_devices (ccmtype, room, region, `order`, `priority`)
VALUES (?, ?, ?, ?, ?);