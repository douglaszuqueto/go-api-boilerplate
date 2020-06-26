package utils

import (
	"net/http"

	"github.com/douglaszuqueto/go-api-boilerplate/pkg/types"
	"github.com/gorilla/mux"
)

// Route route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes routes
type Routes []Route

// Error error
type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Response as
type Response struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

// Data data response
type Data struct {
	Data interface{} `json:"data"`
}

// AddRoute AddRoute
func AddRoute(path string, method string, handler http.HandlerFunc) Route {
	return Route{path, method, handler}
}

// GetVar getVar
func GetVar(r *http.Request, key string) types.UUID {
	vars := mux.Vars(r)
	id := vars[key]

	return types.UUID(id)
}

// GetID getVar
func GetID(r *http.Request) types.UUID {
	vars := mux.Vars(r)
	id := vars["id"]

	return types.UUID(id)
}

// ResponseSuccess ResponseSuccess
func ResponseSuccess(message string, i interface{}) Response {
	return Response{Error{false, message}, i}
}

// ResponseError ResponseError
func ResponseError(message string) Response {
	return Response{Error{true, message}, nil}
}
