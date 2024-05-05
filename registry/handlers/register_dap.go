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
	w.WriteHeader(http.StatusNotImplemented)
}
