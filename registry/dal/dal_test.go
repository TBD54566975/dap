package dal_test

import (
	"context"
	libdal "dapregistry/dal"
	"database/sql"
	"encoding/json"
	"errors"
	"testing"

	"github.com/TBD54566975/dap"
	"github.com/alecthomas/assert"
	"github.com/tbd54566975/web5-go/dids/didjwk"
)

func getDAL(t *testing.T) (*libdal.DAL, func(ctx context.Context)) {
	db, err := sql.Open("sqlite", "../db/registry.sqlite3")

	assert.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
	})

	clearTables := func(ctx context.Context) {
		db.ExecContext(ctx, "DELETE FROM daps")
	}

	return libdal.New(db), clearTables
}

func TestCreateDAP(t *testing.T) {
	dal, clearTables := getDAL(t)
	ctx := context.Background()

	t.Run("works", func(t *testing.T) {
		t.Cleanup(func() {
			clearTables(ctx)
		})

		aliceDID, err := didjwk.Create()
		assert.NoError(t, err)

		reg := dap.NewRegistration("alice", "didpay.me", aliceDID.URI)
		err = reg.Sign(aliceDID)
		assert.NoError(t, err)

		proof, err := json.Marshal(reg)
		assert.NoError(t, err)

		err = dal.CreateDAP(ctx, reg, string(proof))
		assert.NoError(t, err)
	})

	t.Run("same handle", func(t *testing.T) {
		t.Cleanup(func() {
			clearTables(ctx)
		})

		aliceDID, err := didjwk.Create()
		assert.NoError(t, err)

		reg := dap.NewRegistration("alice", "didpay.me", aliceDID.URI)
		err = reg.Sign(aliceDID)
		assert.NoError(t, err)

		proof, err := json.Marshal(reg)
		assert.NoError(t, err)

		err = dal.CreateDAP(ctx, reg, string(proof))
		assert.NoError(t, err)

		bobDID, err := didjwk.Create()
		assert.NoError(t, err)

		reg2 := dap.NewRegistration("alice", "didpay.me", bobDID.URI)
		err = reg2.Sign(bobDID)
		assert.NoError(t, err)

		proof2, err := json.Marshal(reg2)
		assert.NoError(t, err)

		err = dal.CreateDAP(ctx, reg2, string(proof2))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, libdal.ErrHandleConflict))
	})

	t.Run("same did", func(t *testing.T) {
		t.Cleanup(func() {
			clearTables(ctx)
		})

		aliceDID, err := didjwk.Create()
		assert.NoError(t, err)

		reg := dap.NewRegistration("alice", "didpay.me", aliceDID.URI)
		err = reg.Sign(aliceDID)
		assert.NoError(t, err)

		proof, err := json.Marshal(reg)
		assert.NoError(t, err)

		err = dal.CreateDAP(ctx, reg, string(proof))
		assert.NoError(t, err)

		reg2 := dap.NewRegistration("alice2", "didpay.me", aliceDID.URI)
		err = reg2.Sign(aliceDID)
		assert.NoError(t, err)

		proof2, err := json.Marshal(reg2)
		assert.NoError(t, err)

		err = dal.CreateDAP(ctx, reg2, string(proof2))
		assert.Error(t, err)
		assert.True(t, errors.Is(err, libdal.ErrDIDConflict))
	})
}
