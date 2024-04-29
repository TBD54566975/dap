package dap

import "github.com/alecthomas/types/optional"

type ResolutionResponse struct {
	DID   string
	Proof optional.Option[RegistrationRequest]
}
