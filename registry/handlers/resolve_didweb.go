package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

type ResolveDIDWeb struct {
	DIDDocument didcore.Document
}

func (h *ResolveDIDWeb) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: marshal did document and send it back
	w.WriteHeader(http.StatusNotImplemented)
}
