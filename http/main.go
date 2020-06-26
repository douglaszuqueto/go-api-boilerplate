package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/douglaszuqueto/go-api-boilerplate/http/api/admin"
	"github.com/douglaszuqueto/go-api-boilerplate/http/api/auth"
	"github.com/douglaszuqueto/go-api-boilerplate/http/api/client"
	"github.com/douglaszuqueto/go-api-boilerplate/http/middleware"
	"github.com/douglaszuqueto/go-api-boilerplate/http/utils"

	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	authService := auth.NewAuthAPI()
	adminService := admin.NewAdminAPI()
	clientService := client.NewClientAPI()

	authService.Register(
		router.PathPrefix("/api/auth").Subrouter(),
	)

	adminService.Register(
		router.PathPrefix("/api/admin").Subrouter(),
	)

	clientService.Register(
		router.PathPrefix("/api").Subrouter(),
	)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(utils.ResponseError("error 404"))
	})

	return router
}

// Run run
func Run() {

	addr := fmt.Sprintf(":%s", "3000")
	log.Printf("[API] API iniciada: http://127.0.0.1:%s", "3000")

	server := http.Server{
		Addr:         addr,
		Handler:      middleware.CORS(routes()),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatalln(server.ListenAndServe())
}
