-- name: GetDevicesFromHouse :many
SELECT 
    id, house_id, climate_data_id, m304_id, sensor_id, device_name, rly, created_at, updated_at
FROM devices
WHERE house_id = ?;

-- name: GetJoinedDevicesFromHouse :many
SELECT 
    d.id, d.house_id, d.m304_id, d.sensor_id, d.device_name, d.rly, d.created_at, d.updated_at,
    c.name AS climate_data_name, c.unit
FROM devices d
JOIN climate_datas c ON d.climate_data_id = c.id
WHERE d.house_id = ?;

-- name: CreateDevice :execlastid
INSERT INTO devices (house_id, climate_data_id, m304_id, sensor_id, device_name, rly) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetDeviceFromID :one
SELECT id, house_id, climate_data_id, m304_id, sensor_id, device_name, rly
FROM devices
WHERE id = ?;