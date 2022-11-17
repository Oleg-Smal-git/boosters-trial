package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ServeOK serves a 200 response.
func ServeOK(writer http.ResponseWriter, request *http.Request, object interface{}) {
	writer.WriteHeader(http.StatusOK)
	RenderJSON(writer, request, object)
}

// ServeEmptyOK serves an empty 200 response.
func ServeEmptyOK(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

// ServeCreated serves a 201 response with a provided object.
func ServeCreated(writer http.ResponseWriter, request *http.Request, object interface{}) {
	writer.WriteHeader(http.StatusCreated)
	RenderJSON(writer, request, object)
}

// ServeBadRequest serves a 400 response with a provided message.
func ServeBadRequest(writer http.ResponseWriter, request *http.Request, message string) {
	writer.WriteHeader(http.StatusBadRequest)
	response := make(map[string]string)
	response["message"] = message
	RenderJSON(writer, request, response)
}

// ServeNotFound serves a standard 404 error.
func ServeNotFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	response := make(map[string]string)
	response["message"] = fmt.Sprintf("no handler registered at route %v for method %v", request.Method, request.URL.Path)
	RenderJSON(writer, request, response)
}

// ServeConflict serves a 409 response with a provided message.
func ServeConflict(writer http.ResponseWriter, request *http.Request, message string) {
	writer.WriteHeader(http.StatusConflict)
	response := make(map[string]string)
	response["message"] = message
	RenderJSON(writer, request, response)
}

// RenderJSON writes a json to response.
func RenderJSON(writer http.ResponseWriter, request *http.Request, response interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	bytesResponse, err := json.Marshal(response)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to marshal response")
	}
	_, err = writer.Write(bytesResponse)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to write response")
	}
}
