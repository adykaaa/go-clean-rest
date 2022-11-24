package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct{}

var chiDispatcher = chi.NewRouter()

func NewChiRouter() Router {
	return chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*chiRouter) Serve(port string) {
	fmt.Printf("Server listening on port %s", port)
	http.ListenAndServe(port, muxDispatcher)
}
