package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adykaaa/go-clean-rest/controller"
	router "github.com/adykaaa/go-clean-rest/http"
	"github.com/adykaaa/go-clean-rest/repository"
	"github.com/adykaaa/go-clean-rest/service"
)

var (

	/*
		The beauty in abstractions using interfaces is this:it literally takes 1-1 variable
				replacements to use a completely different Router library and Database to store stuff. Hell yeah!

				postRepository repository.PostRepository = repository.NewSQLRepository()
				httpRouter     router.Router             = router.NewChiRouter()
	*/
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
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
