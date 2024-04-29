package maddr

import (
	"fmt"

	"github.com/TBD54566975/dap/maddr/urn"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

const (
	MoneyAddrKind = "maddr"
)

type MoneyAddress struct {
	ID       string
	URN      urn.URN
	Currency string
	CSS      string
}

func FromDIDService(svc didcore.Service) ([]MoneyAddress, error) {
	if svc.Type != MoneyAddrKind {
		return nil, fmt.Errorf("invalid service type: %s", svc.Type)
	}

	maddrs := make([]MoneyAddress, len(svc.ServiceEndpoint))

	for i, se := range svc.ServiceEndpoint {
		urn, err := urn.Decode(se)
		if err != nil {
			return nil, fmt.Errorf("invalid money address: %w", err)
		}

		maddr := MoneyAddress{
			URN:      urn,
			ID:       svc.ID,
			Currency: urn.NID,
			CSS:      urn.NSS,
		}

		maddrs[i] = maddr
	}

	return maddrs, nil
}
