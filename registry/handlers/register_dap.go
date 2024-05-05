package handlers

import (
	"dapregistry/dal"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/julienschmidt/httprouter"
)

type RegisterDAP struct {
	DAL *dal.DAL
	Log logr.Logger
}

func (h *RegisterDAP) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/* TODO:
	 * unmarshal request body into RegistrationRequest
	 * check to ensure provided did's method is supported
	 * verify signature
	 * write to db
	 * return 409 conflict if handle is already taken
	 * sign registration request digest
	 * construct RegistrationResponse
	 * return 202 accepted
	**/

	w.WriteHeader(http.StatusNotImplemented)
}
