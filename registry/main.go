package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func CORS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "" {
		header := w.Header()

		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
	}

	w.WriteHeader(http.StatusNoContent)
}

func Challenge(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNotImplemented)
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNotImplemented)
}

func ResolveHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNotImplemented)
}

func Metadata(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNotImplemented)
}

func ResolveDIDWeb(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNotImplemented)
}

func main() {

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(CORS)

	router.GET("/challenge", Challenge)
	router.GET("/metadata", Metadata)
	router.GET("/resolve/:handle", ResolveHandle)
	router.GET("/.well-known/did.json", ResolveDIDWeb)

	router.POST("/register", Register)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	err := http.ListenAndServe(":8080", n)
	if err != nil {
		log.Fatal(err)
	}
}
