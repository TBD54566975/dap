package dap

import (
	"context"
	"fmt"

	"github.com/TBD54566975/dap/maddr"
	"github.com/alecthomas/types/optional"
	"github.com/tbd54566975/web5-go/dids"
)

// Resolve resolves a DAP to a set of money addresses
func Resolve(dap string) ([]maddr.MoneyAddress, error) {
	return ResolveWithContext(context.Background(), dap)
}

// ResolveWithContext resolves a DAP to a set of money addresses
func ResolveWithContext(ctx context.Context, dap string) ([]maddr.MoneyAddress, error) {
	resp, err := client.Resolve(ctx, dap)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve dap: %w", err)
	}

	// TODO: check to make sure resp can't be nil
	result, err := dids.ResolveWithContext(ctx, resp.DID)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve DID %s: %w", resp.DID, err)
	}

	maddrs := make([]maddr.MoneyAddress, 0)
	for _, svc := range result.Document.Service {
		m, err := maddr.FromDIDService(svc)
		if err != nil {
			return nil, fmt.Errorf("failed to parse money address in service %s: %w", svc.ID, err)
		}

		maddrs = append(maddrs, m...)
	}

	return maddrs, err
}

type ResolutionResponse struct {
	DID   string                               `json:"did,omitempty"`
	Proof optional.Option[RegistrationRequest] `json:"proof"`
}
