package dap

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	ServiceType string = "dapregistry"
)

// TODO: remove this eventually
var client = NewClient()

func Parse(input string) (*DAP, error) {
	delimIdx := strings.LastIndex(input, "@")
	if delimIdx == -1 {
		return nil, errors.New("expected format '<handle>@domain'")
	}

	handle := input[:delimIdx]
	domain := input[delimIdx+1:]

	if len(domain) == 0 {
		return nil, errors.New("domain cannot be empty")
	}

	if len(handle) < 3 || len(handle) > 30 {
		return nil, errors.New("handle must be between 3-30 characters")
	}

	for i, c := range handle {
		if unicode.IsControl(c) {
			return nil, fmt.Errorf("invalid character in handler at pos %d", i)
		} else if unicode.IsPunct(c) {
			return nil, fmt.Errorf("invalid character in handler at pos %d", i)
		}
	}

	return &DAP{Handle: handle, Domain: domain}, nil
}

type DAP struct {
	Handle string
	Domain string
}

func (d DAP) String() string {
	return fmt.Sprintf("%s@%s", d.Handle, d.Domain)
}

func (d DAP) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DAP) UnmarshalText(text []byte) error {
	s := string(text)
	parsedDAP, err := Parse(s)
	if err != nil {
		return err
	}

	*d = *parsedDAP
	return nil
}
