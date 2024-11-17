-- name: GetTimeSchedulesFromDeviceCondition :many
SELECT id, device_condition_id, start_time, end_time
FROM time_schedules
WHERE device_condition_id = ?;

-- name: GetTimeScheduleFromID :one
SELECT id, device_condition_id, start_time, end_time
FROM time_schedules
WHERE id = ?;

-- name: CreateTimeSchedule :execlastid
INSERT INTO time_schedules (device_condition_id, start_time, end_time)
VALUES (?, ?, ?);