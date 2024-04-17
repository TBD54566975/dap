// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
	"time"
)

const createDAP = `-- name: CreateDAP :exec
INSERT INTO daps (id, did, handle, date_created)
VALUES (?, ?, ?, ?)
`

type CreateDAPParams struct {
	ID          string
	Did         string
	Handle      string
	DateCreated time.Time
}

func (q *Queries) CreateDAP(ctx context.Context, arg CreateDAPParams) error {
	_, err := q.db.ExecContext(ctx, createDAP,
		arg.ID,
		arg.Did,
		arg.Handle,
		arg.DateCreated,
	)
	return err
}