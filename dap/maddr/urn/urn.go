package urn

import (
	"fmt"
	"strings"
)

func Parse(input string) (URN, error) {
	if !strings.HasPrefix(input, "urn:") {
		return URN{}, fmt.Errorf("expected urn. got %s", input)
	}

	urnless := input[4:]

	delimIDX := strings.IndexRune(urnless, ':')
	if delimIDX == -1 {
		return URN{}, fmt.Errorf("invalid money address. expected urn:[currency]:[css]. got %s", input)
	}

	return URN{
		URN: input,
		NID: urnless[:delimIDX],
		NSS: urnless[delimIDX+1:],
	}, nil
}

type URN struct {
	URN string
	NID string
	NSS string
}
