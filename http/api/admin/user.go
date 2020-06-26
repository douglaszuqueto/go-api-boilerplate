package admin

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

type storeRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	State    uint32 `json:"state"`
}

type updateRequest struct {
	Password string `json:"-"`
	State    uint32 `json:"state"`
}

// Store store
func (h *userAPI) Store(w http.ResponseWriter, r *http.Request) {
	var request storeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	user := storage.User{
		Username:  request.Username,
		Password:  request.Password,
		State:     request.State,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := storage.GetDB().CreateUser(context.Background(), user)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess(id, nil))
}

// Update update
func (h *userAPI) Update(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	if err := id.IsValid(); err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	userDB, err := storage.GetDB().GetUser(context.Background(), id.String())
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	var request updateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	user := storage.User{
		ID:        userDB.ID,
		Username:  userDB.Username,
		Password:  request.Password,
		State:     request.State,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = storage.GetDB().UpdateUser(context.Background(), user)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess(user.ID, nil))
}

// Remove remove
func (h *userAPI) Remove(w http.ResponseWriter, r *http.Request) {
	id := utils.GetID(r)

	if err := id.IsValid(); err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	_, err := storage.GetDB().GetUser(context.Background(), id.String())
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	err = storage.GetDB().DeleteUser(context.Background(), id.String())
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess("", nil))
}
