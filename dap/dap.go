package dap

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	ServiceType string = "dapregistry"
)

func Parse(input string) (*DAP, error) {
	// Regex to validate the full DAP format: local-handle@domain
	dapRegex := regexp.MustCompile(`^([^;!@%^&*()/\\]{3,30})@(.+)$`)
	matches := dapRegex.FindStringSubmatch(input)

	fmt.Println(len(matches))

	if matches == nil {
		return nil, errors.New("invalid DAP format")
	}

	handle, domain := matches[1], matches[2]

	// Check if the domain part is not empty (additional domain validations can be added as needed)
	if domain == "" {
		return nil, errors.New("domain cannot be empty")
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
