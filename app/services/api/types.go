package api

import (
	"net/http"
)

// Serve is an alias to a http endpoint functional handler.
type Serve func(http.ResponseWriter, *http.Request)

// handler is a wrapper for Serve that implements http.Handler and is used purely for adaptation purposes.
type handler struct {
	serve Serve
}

// ServeHTTP wraps Serve.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.serve(w, r)
}
