package admin

import (
	"github.com/douglaszuqueto/go-api-boilerplate/http/middleware"
	"github.com/douglaszuqueto/go-api-boilerplate/http/utils"

	"github.com/gorilla/mux"
)

// Service service
type Service struct {
	user *userAPI
}

// NewAdminAPI NewAdminAPI
func NewAdminAPI() *Service {
	return &Service{
		user: newUserAPI(),
	}
}

// Register Register
func (s *Service) Register(router *mux.Router) {
	router.Use(middleware.MaxClients)
	router.Use(middleware.Logger)

	routes := utils.Routes{

		// User
		utils.AddRoute("/user", "GET", s.user.All),
		utils.AddRoute("/user", "POST", s.user.Store),
		utils.AddRoute("/user/{id}", "GET", s.user.Get),
		utils.AddRoute("/user/{id}", "PUT", s.user.Update),
		utils.AddRoute("/user/{id}", "DELETE", s.user.Remove),
	}

	for _, r := range routes {
		router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}
