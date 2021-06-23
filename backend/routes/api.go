package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Api(r *mux.Router) {
	r.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("cc"))
	})
}
