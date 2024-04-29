package dap_test

import (
	"testing"

	"github.com/TBD54566975/dap"
	"github.com/alecthomas/assert"
)

func TestParse(t *testing.T) {
	vectors := []struct {
		input    string
		expected dap.DAP
		err      bool
	}{
		{
			input:    "moegrammer@didpay.me",
			expected: dap.DAP{Handle: "moegrammer", Domain: "didpay.me"},
		},
		{
			input:    "moegrammer@www.linkedin.com",
			expected: dap.DAP{Handle: "moegrammer", Domain: "www.linkedin.com"},
		},
		{
			input: "doodoo",
			err:   true,
		},
		{
			input: "doo@doo@doodoo.gov",
			err:   true,
		},
		{
			input: "doodoo@",
			err:   true,
		},
		{
			input: "@",
			err:   true,
		},
	}

	for _, v := range vectors {
		t.Run(v.input, func(t *testing.T) {
			actual, err := dap.Parse(v.input)
			if v.err {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, actual)

				assert.Equal(t, v.expected.Handle, actual.Handle)
				assert.Equal(t, v.expected.Domain, actual.Domain)
			}
		})
	}
}
