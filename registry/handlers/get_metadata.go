package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: include metadata options in struct
type GetMetadata struct{}

func (h *GetMetadata) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: marshal metadata options and send them back
	w.WriteHeader(http.StatusNotImplemented)
}
