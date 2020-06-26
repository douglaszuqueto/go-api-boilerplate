package client

import (
	"github.com/douglaszuqueto/go-api-boilerplate/http/middleware"
	"github.com/douglaszuqueto/go-api-boilerplate/http/utils"

	"github.com/gorilla/mux"
)

// Service service
type Service struct {
	user *userAPI
}

// NewClientAPI NewClientAPI
func NewClientAPI() *Service {
	return &Service{
		user: newUserAPI(),
	}
}

// Register Register
func (s *Service) Register(router *mux.Router) {
	router.Use(middleware.Logger)
	router.Use(middleware.MaxClients)

	routes := utils.Routes{

		// User
		utils.AddRoute("/user", "GET", s.user.All),
		utils.AddRoute("/user/{id}", "GET", s.user.Get),
	}

	for _, r := range routes {
		router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}
