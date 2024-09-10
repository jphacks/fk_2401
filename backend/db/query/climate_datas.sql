-- name: GetAllClimateData :many
SELECT id, name, unit
FROM climate_datas;

-- name: GetClimateDataFromID :one
SELECT id, name, unit
FROM climate_datas
WHERE id = ?;
