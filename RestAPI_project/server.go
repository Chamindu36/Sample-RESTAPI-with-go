package main

import (
	"./contoller"
	"./router"
	"./service"
	"./repository"
	"fmt"
	"net/http"
)

var (
	postRepository repository.PostRepository = repository.NewFireStoreRepository()
	postService service.PostService = service.NewPostService(postRepository)
	postController contoller.PostController = contoller.NewPostController(postService)
	httpRouter     router.Router            = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and Running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/addPost", postController.AddPost)
	httpRouter.GET("/getPost", postController.GetPost)
	httpRouter.Serve(port)
}
