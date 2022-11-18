package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Oleg-Smal-git/boosters-trial/app/services/api"
	"github.com/Oleg-Smal-git/boosters-trial/app/services/database"
	eposts "github.com/Oleg-Smal-git/boosters-trial/app/services/posts/entities"
	iposts "github.com/Oleg-Smal-git/boosters-trial/app/services/posts/interfaces"
	posts "github.com/Oleg-Smal-git/boosters-trial/app/services/posts/logic"

	"github.com/gorilla/mux"
)

var (
	postsService iposts.PostsService
)

// PostsController is a wrapper for controllers that interact with posts.
type PostsController struct{}

// MustInitialize performs all the setup needed for the controller.
func (PostsController) MustInitialize() {
	ctx := context.Background()
	reader, err := database.GetReader(ctx)
	if err != nil {
		panic(err)
	}
	writer, err := database.GetWriter(ctx)
	if err != nil {
		panic(err)
	}
	service, err := posts.NewPostsService(ctx, reader, writer)
	if err != nil {
		panic(err)
	}
	postsService = service
}

// IndexPosts fetches all posts.
func (PostsController) IndexPosts(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	ps, err := postsService.IndexPosts(ctx)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	api.ServeOK(writer, request, ps)
}

// FindPost fetches a single post.
func (PostsController) FindPost(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	rawID := mux.Vars(request)["id"]
	id, err := strconv.Atoi(rawID)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	p, err := postsService.FindPost(ctx, eposts.PostID(id))
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	api.ServeOK(writer, request, p)
}

// UpdatePost updates a post.
func (PostsController) UpdatePost(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	rawID := mux.Vars(request)["id"]
	id, err := strconv.Atoi(rawID)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	var p eposts.Post
	err = api.ParseJSONBody(&p, request.Body)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	p.ID = eposts.PostID(id)
	err = postsService.UpdatePost(ctx, &p)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	api.ServeCreated(writer, request, p)
}

// CreatePost creates a post.
func (PostsController) CreatePost(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	var p eposts.Post
	err := api.ParseJSONBody(&p, request.Body)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	err = postsService.CreatePost(ctx, &p)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	api.ServeCreated(writer, request, p)
}

// DeletePost deletes a post.
func (PostsController) DeletePost(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	rawID := mux.Vars(request)["id"]
	id, err := strconv.Atoi(rawID)
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	err = postsService.DeletePost(ctx, eposts.PostID(id))
	if err != nil {
		api.ServeBadRequest(writer, request, err.Error())
		return
	}
	api.ServeMessageOK(writer, request, "post deleted")
}
