package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/douglaszuqueto/go-api-boilerplate/http/utils"
	"github.com/douglaszuqueto/go-api-boilerplate/pkg/storage"
)

type userAPI struct{}

func newUserAPI() *userAPI {
	return new(userAPI)
}

// All all
func (h *userAPI) All(w http.ResponseWriter, r *http.Request) {
	users, err := storage.GetDB().ListUser(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess("", users))
}

// Get get
func (h *userAPI) Get(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	if err := id.IsValid(); err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	user, err := storage.GetDB().GetUser(context.Background(), id.String())
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess("", user))
}
