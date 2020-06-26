package auth

import (
	"github.com/douglaszuqueto/go-api-boilerplate/http/middleware"
	"github.com/gorilla/mux"
)

// Service service
type Service struct {
	login *loginAPI
}

// NewAuthAPI NewAuthAPI
func NewAuthAPI() *Service {
	return &Service{
		login: newLoginAPI(),
	}
}

// Register Register
func (s *Service) Register(router *mux.Router) {
	router.Use(middleware.Logger)
	router.Use(middleware.MaxClients)

	router.HandleFunc("/admin/signin", s.login.AdminAuth).Methods("POST")
	router.HandleFunc("/client/signin", s.login.ClientAuth).Methods("POST")
}
