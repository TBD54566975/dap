package dal

import (
	"context"
	"dapregistry/dal/sqlc"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/TBD54566975/dap"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"

	"github.com/alecthomas/types/optional"
)

var ErrDIDConflict = errors.New("did already registered")
var ErrHandleConflict = errors.New("handle already registered")

type DAL struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func New(db *sql.DB) *DAL {
	return &DAL{
		db:      db,
		queries: sqlc.New(db),
	}
}

type HandleDID struct {
	DID   string
	Proof optional.Option[dap.RegistrationRequest]
}

func (d *DAL) GetHandleDID(ctx context.Context, handle string) (*HandleDID, error) {
	entry, err := d.queries.GetHandleDID(ctx, handle)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("querying failed: %w", err)
	}

	var proof optional.Option[dap.RegistrationRequest]
	if p, ok := entry.Proof.Get(); ok {
		var registrationRequest dap.RegistrationRequest

		err = json.Unmarshal([]byte(p), &registrationRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal proof: %w", err)
		}

		proof = optional.Some(registrationRequest)
	} else {
		proof = optional.None[dap.RegistrationRequest]()
	}

	return &HandleDID{
		DID:   entry.Did,
		Proof: proof,
	}, nil
}

func (d *DAL) CreateDAP(ctx context.Context, r dap.RegistrationRequest, proof string) error {
	params := sqlc.CreateDAPParams{
		ID:          r.ID.String(),
		Did:         r.DID,
		Handle:      r.Handle,
		Proof:       optional.Some(proof),
		DateCreated: time.Now().UTC().Format(time.RFC3339),
	}

	err := d.queries.CreateDAP(ctx, params)
	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			msg := sqlErr.Error()
			if strings.Contains(msg, "handle") {
				return ErrHandleConflict
			} else if strings.Contains(msg, "did") {
				return ErrDIDConflict
			}
		}

		return err
	}

	return err
}
