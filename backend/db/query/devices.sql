-- name: GetDevicesFromHouse :many
SELECT 
    id, house_id, climate_data_id, uecs_device_id, device_name, valid, set_point, duration, operator, created_at, updated_at
FROM devices
WHERE house_id = ?;

-- name: GetJoinedDevicesFromHouse :many
SELECT 
    d.id, d.house_id, d.uecs_device_id, d.device_name, d.valid, d.set_point, d.duration, d.operator, d.created_at, d.updated_at,
    c.name AS climate_data_name, c.unit
FROM devices d
JOIN climate_datas c ON d.climate_data_id = c.id
WHERE d.house_id = ?;

-- name: CreateDevice :execlastid
INSERT INTO devices (house_id, climate_data_id, uecs_device_id, device_name, valid, set_point, duration, operator) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?);
