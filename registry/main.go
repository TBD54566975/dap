package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"dapregistry/dal"
	"dapregistry/handlers"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

func main() {
	var log logr.Logger

	zapLog, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("failed to create logger %v\n", err)
		os.Exit(1)
	}

	log = zapr.NewLogger(zapLog)

	db, err := sql.Open("sqlite", "db/registry.sqlite3")
	if err != nil {
		log.Error(err, "failed to open database")
		os.Exit(1)
	}

	dal := dal.New(db)
	defer db.Close()

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(CORS)

	getMetadata := handlers.GetMetadata{
		Metadata: handlers.Metadata{
			Registration: handlers.RegistrationMetadata{
				Enabled:             true,
				SupportedDIDMethods: []string{"dht"},
			},
		},
	}
	router.GET("/metadata", getMetadata.Handle)

	resolveDIDWeb := handlers.ResolveDIDWeb{} // TODO: set DIDDocument field
	router.GET("/.well-known/did.json", resolveDIDWeb.Handle)

	resolveDAP := handlers.ResolveDAP{DAL: dal, Log: log}
	router.GET("/daps/:handle", resolveDAP.Handle)

	registerDAP := handlers.RegisterDAP{DAL: dal, Log: log}
	router.POST("/daps", registerDAP.Handle)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	log.Info("server listening on 8080")
	err = http.ListenAndServe(":8080", n)
	if err != nil {
		log.Error(err, "failed to start server")
		os.Exit(1)
	}
}

func CORS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "" {
		header := w.Header()

		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
	}

	w.WriteHeader(http.StatusNoContent)
}
