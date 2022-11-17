package controllers

var (
	// These are controller instances.
	postsController	= PostsController{}
)

var (
	// handlers is a map of routes and functions that control them.
	handlers = map[string]map[string]string { 
		"DELETE": {
			"/posts/:id": postsController.DeletePost,
		},
		"GET": {
			"/posts": postsController.IndexPosts,"/posts/:id": postsController.FindPost,
		},
		"POST": {
			"/posts": postsController.CreatePost,
		},
		"PUT": {
			"/posts/:id": postsController.UpdatePost,
		},
	}
)