package dap

import "github.com/alecthomas/types/optional"

type ResolutionResponse struct {
	DID   string                               `json:"did,omitempty"`
	Proof optional.Option[RegistrationRequest] `json:"proof"`
}
