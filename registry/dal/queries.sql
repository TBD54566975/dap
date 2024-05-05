-- name: CreateDAP :exec
INSERT INTO daps (id, did, handle, date_created)
  VALUES (?, ?, ?, ?);

-- name: GetHandleDID :one
SELECT * FROM daps
  WHERE handle = ?
  LIMIT 1;