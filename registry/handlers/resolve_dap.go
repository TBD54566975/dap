package handlers

import (
	"dapregistry/dal"
	"encoding/json"
	"net/http"

	libdap "github.com/TBD54566975/dap"

	"github.com/go-logr/logr"
	"github.com/julienschmidt/httprouter"
)

type ResolveDAP struct {
	DAL *dal.DAL
	Log logr.Logger
}

func (h *ResolveDAP) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	handle := p.ByName("handle")
	dap, err := libdap.Parse(handle + "@" + r.Host)
	if err != nil {
		body := libdap.ErrHTTPResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}

		bytes, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(body.Status)
		w.Write(bytes)

		return
	}

	handleDID, err := h.DAL.GetHandleDID(r.Context(), dap.Handle)
	if err != nil {
		h.Log.Error(err, "failed to get dap")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if handleDID == nil {
		body := libdap.ErrHTTPResponse{
			Status:  http.StatusNotFound,
			Message: "Not Found",
		}

		bytes, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(body.Status)
		w.Write(bytes)

		return
	}

	body := libdap.ResolutionResponse{
		DID:   handleDID.DID,
		Proof: handleDID.Proof,
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
