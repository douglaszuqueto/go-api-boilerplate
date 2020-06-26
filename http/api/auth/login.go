package auth

import (
	"encoding/json"
	"net/http"

	"github.com/douglaszuqueto/go-api-boilerplate/http/utils"
)

type loginAPI struct{}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func newLoginAPI() *loginAPI {
	return new(loginAPI)
}

// AdminAuth AdminAuth
func (h *loginAPI) AdminAuth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", nil))
}

// ClientSignIn ClientSignIn
func (h *loginAPI) ClientAuth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", nil))
}
