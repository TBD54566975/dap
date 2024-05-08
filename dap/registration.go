package dap

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/gowebpki/jcs"
	"github.com/tbd54566975/web5-go/dids/did"
	"github.com/tbd54566975/web5-go/jws"
	"go.jetpack.io/typeid"
)

type RegistrationID struct {
	typeid.TypeID[RegistrationID]
}

func (r RegistrationID) Prefix() string { return "reg" }

type RegistrationRequest struct {
	ID        RegistrationID `json:"id"`
	Handle    string         `json:"handle"`
	DID       string         `json:"did"`
	Domain    string         `json:"domain"`
	Signature string         `json:"signature"`
}

func NewRegistration(handle, domain, did string) RegistrationRequest {
	return RegistrationRequest{
		ID:     typeid.Must(typeid.New[RegistrationID]()),
		Handle: handle,
		DID:    did,
		Domain: domain,
	}
}

func (r RegistrationRequest) Digest() ([]byte, error) {
	payload := map[string]string{
		"id":     r.ID.String(),
		"handle": r.Handle,
		"domain": r.Domain,
		"did":    r.DID,
	}

	m, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	canonicalized, err := jcs.Transform(m)
	if err != nil {
		return nil, fmt.Errorf("failed to canonicalize payload: %w", err)
	}

	checksum := sha256.Sum256(canonicalized)
	return checksum[:], nil

}

func (r *RegistrationRequest) Sign(bearerDID did.BearerDID) error {
	digest, err := r.Digest()
	if err != nil {
		return fmt.Errorf("failed to compute digest: %w", err)
	}

	compactJWS, err := jws.Sign(digest, bearerDID, jws.DetachedPayload(true))
	if err != nil {
		return fmt.Errorf("failed to compute signature: %w", err)
	}

	r.Signature = compactJWS

	return nil
}

func (r RegistrationRequest) Verify() (jws.Decoded, error) {
	digest, err := r.Digest()
	if err != nil {
		return jws.Decoded{}, fmt.Errorf("failed to compute digest: %w", err)
	}

	decoded, err := jws.Verify(r.Signature, jws.Payload(digest))
	if err != nil {
		return jws.Decoded{}, fmt.Errorf("failed to verify signature: %w", err)
	}

	return decoded, nil
}

type RegistrationResponse struct{}

// Register registers the provided DID with the provided DAP at the DAP's respective registry
func Register(dap string, bearerDID did.BearerDID) error {
	return RegisterWithContext(context.Background(), dap, bearerDID)
}

func RegisterWithContext(ctx context.Context, dap string, bearerDID did.BearerDID) error {
	d, err := Parse(dap)
	if err != nil {
		return fmt.Errorf("failed to parse dap: %w", err)
	}

	// TODO: change last arg to type bearerDID
	req := NewRegistration(d.Handle, d.Domain, bearerDID.URI)
	if err := req.Sign(bearerDID); err != nil {
		return fmt.Errorf("failed to sign registration request: %w", err)
	}

	_, err = client.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	return nil
}
