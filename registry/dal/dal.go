package dal

import (
	"context"
	"dapregistry/dal/sqlc"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TBD54566975/dap"

	"github.com/alecthomas/types/optional"
)

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
