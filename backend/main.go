package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oxodao/cardsagainstoverflow/routes"
)

const (
	CAO_VERSION = "3.0-indev"
	CAO_AUTHOR  = "Oxodao"
)

//go:embed web
var frontend embed.FS

func main() {
	fmt.Printf("CardsAgainstOverflow v%v - by %v\n", CAO_VERSION, CAO_AUTHOR)

	realFrontend, _ := fs.Sub(frontend, "web")

	r := mux.NewRouter()
	routes.Api(r.PathPrefix("/api").Subrouter())
	r.PathPrefix("/").Handler(http.FileServer(http.FS(realFrontend)))

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
