package handlers

import (
	"dapregistry/dal"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/TBD54566975/dap"
	"github.com/go-logr/logr"
	"github.com/julienschmidt/httprouter"
)

type RegisterDAP struct {
	DAL *dal.DAL
	Log logr.Logger
}

func (h *RegisterDAP) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(body) == 0 {
		errResp := dap.ErrHTTPResponse{Status: http.StatusBadRequest, Message: "expected request body"}
		w.WriteHeader(errResp.Status)

		if resp, err := json.Marshal(errResp); err == nil {
			w.Write(resp)
		} else {
			h.Log.Error(err, "failed to marshal HTTP Response")
		}

		return
	}

	var registrationRequest dap.RegistrationRequest
	if err := json.Unmarshal(body, &registrationRequest); err != nil {
		errResp := dap.ErrHTTPResponse{Status: http.StatusBadRequest, Message: "expected valid request body"}
		w.WriteHeader(errResp.Status)

		if resp, err := json.Marshal(errResp); err == nil {
			w.Write(resp)
		} else {
			h.Log.Error(err, "failed to marshal HTTP Response")
		}

		return
	}

	// TODO: check to ensure provided DID's method is supported

	decodedJWS, err := registrationRequest.Verify()
	if err != nil {
		// TODO: fill out
	}

	if decodedJWS.SignerDID.String() != registrationRequest.DID {
		// TODO: return 400 - signer does not match did in request body
	}

	err = h.DAL.CreateDAP(r.Context(), registrationRequest, string(body))
	if err != nil {
		if errors.Is(err, dal.ErrHandleConflict) {
			respErr := dap.ErrHTTPResponse{Status: http.StatusConflict, Message: "Handle taken"}
			w.WriteHeader(respErr.Status)

			if resp, err := json.Marshal(respErr); err == nil {
				w.Write(resp)
			} else {
				h.Log.Error(err, "failed to marshal response")
			}
		} else if errors.Is(err, dal.ErrDIDConflict) {
			respErr := dap.ErrHTTPResponse{Status: http.StatusConflict, Message: "DID already registered"}
			w.WriteHeader(respErr.Status)

			if resp, err := json.Marshal(respErr); err == nil {
				w.Write(resp)
			} else {
				h.Log.Error(err, "failed to marshal response")
			}
		} else {
			h.Log.Error(err, "failed to create DAP")
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusAccepted)

	// TODO: Sign registration request digest and include in response
	resp := dap.RegistrationResponse{}
	if m, err := json.Marshal(resp); err == nil {
		w.Write(m)
	}

	return
}
