-- name: GetRecordFromM304ID :many
SELECT id, m304_id, device_id, `block`, valid, position
FROM m304_record
WHERE m304_id = ?;

-- name: CreateM304Record :execlastid
INSERT INTO m304_record (m304_id, device_id, `block`, valid, position)
VALUES (?, ?, ?, ?, ?);