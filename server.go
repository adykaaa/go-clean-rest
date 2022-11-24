package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adykaaa/go-clean-rest/controller"
	router "github.com/adykaaa/go-clean-rest/http"
)

var (

	/*
		The beauty in abstracting the Router away is this: it literally takes 1 variable
		replacement to use a completely different Router library. Hell yeah!

		httpRouter     router.Router             = router.NewChiRouter()
	*/
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

func main() {
	//temporary location for Firestore creds
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/adykaaa/Downloads/pragmatic-reviews-428f7-firebase-adminsdk-6wmp3-5a923f8d33.json")

	httpRouter.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Server is up and running...")
	})

	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.GET("/posts", postController.GetPosts)

	httpRouter.Serve(":8000")
}
