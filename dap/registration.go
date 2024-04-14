package dap

import (
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

type Registration struct {
	ID        RegistrationID `json:"id"`
	Handle    string         `json:"handle"`
	DID       string         `json:"did"`
	Domain    string         `json:"domain"`
	Signature string         `json:"signature"`
}

func NewRegistration(handle, domain, did string) Registration {
	return Registration{
		ID:     typeid.Must(typeid.New[RegistrationID]()),
		Handle: handle,
		DID:    did,
		Domain: domain,
	}
}

func (r Registration) Digest() ([]byte, error) {
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

func (r *Registration) Sign(bearerDID did.BearerDID) error {
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

func (r Registration) Verify() (jws.Decoded, error) {
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
