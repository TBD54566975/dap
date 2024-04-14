-- name: CreateDAP :exec
INSERT INTO daps (id, did, handle, date_created)
VALUES (?, ?, ?, ?);