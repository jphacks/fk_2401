-- name: GetAllHouses :many
SELECT * FROM houses;

-- name: CreateHouse :execlastid
INSERT INTO houses (name) 
VALUES (?);
