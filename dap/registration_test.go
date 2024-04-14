package dap_test

import (
	"testing"

	"github.com/TBD54566975/dap"

	"github.com/alecthomas/assert"
	"github.com/tbd54566975/web5-go/dids/diddht"
)

func TestRegistration(t *testing.T) {
	bearerDID, err := diddht.Create()
	if err != nil {
		assert.NoError(t, err)
	}

	r := dap.NewRegistration("moegrammer", "didpay.me", bearerDID.URI)
	r.Sign(bearerDID)

	_, err = r.Verify()
	assert.NoError(t, err)
}
