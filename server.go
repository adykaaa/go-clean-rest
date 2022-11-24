package main

import (
	"fmt"
	"net/http"

	router "github.com/adykaaa/go-clean-rest/http"
)

var httpRouter router.Router = router.NewMuxRouter()

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})

	/*	httpRouter.POST("/posts", GetPosts).Methods("GET")
		httpRouter.GET("/posts", AddPost).Methods("POST") */

	httpRouter.Serve(port)
}
