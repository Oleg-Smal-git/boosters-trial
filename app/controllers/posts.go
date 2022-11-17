package controllers

import (
	"net/http"
)

// PostsController is a wrapper for controllers that interact with posts.
type PostsController struct{}

// IndexPosts fetches all posts.
func (PostsController) IndexPosts(http.ResponseWriter, *http.Request) {

}

// FindPost fetches a single post.
func (PostsController) FindPost(http.ResponseWriter, *http.Request) {

}

// UpdatePost updates a post.
func (PostsController) UpdatePost(http.ResponseWriter, *http.Request) {

}

// CreatePost creates a post.
func (PostsController) CreatePost(http.ResponseWriter, *http.Request) {

}

// DeletePost deletes a post.
func (PostsController) DeletePost(http.ResponseWriter, *http.Request) {

}