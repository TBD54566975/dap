package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: Move this to dap lib
type Metadata struct {
	Registration RegistrationMetadata `json:"registration"`
}

type RegistrationMetadata struct {
	Enabled             bool     `json:"enabled"`
	SupportedDIDMethods []string `json:"supportedDidMethods,omitempty"`
}
type GetMetadata struct {
	Metadata Metadata
}

func (h *GetMetadata) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bytes, err := json.Marshal(h.Metadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

	return
}
