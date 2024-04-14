package maddr

import (
	"fmt"

	"github.com/TBD54566975/maddr/urn"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

const (
	MoneyAddrKind = "maddr"
)

// TODO: reconsider using the word Money preceeding 'Address'
type MoneyAddress struct {
	ID       string
	URN      urn.URN
	Currency string // TODO: reconsider using the term 'Currency'
	CSS      string
}

func FromDIDService(svc didcore.Service) (MoneyAddress, error) {
	urn, err := urn.Decode(svc.ServiceEndpoint)
	if err != nil {
		return MoneyAddress{}, fmt.Errorf("invalid money address: %w", err)
	}

	maddr := MoneyAddress{
		URN:      urn,
		ID:       svc.ID,
		Currency: urn.NID,
		CSS:      urn.NSS,
	}

	return maddr, nil
}
