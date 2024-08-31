-- name: GetDevicesFromHouse :many
SELECT 
    d.id, d.house_id, d.duration, d.created_at, d.updated_at,
    c.name, c.unit
FROM devices d
JOIN climate_datas c ON d.climate_data_id = c.id
WHERE d.house_id = ?;

-- name: CreateDevice :execlastid
INSERT INTO devices (house_id, climate_data_id, duration) 
VALUES (?, ?, ?);
