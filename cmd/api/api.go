package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pts/mdes/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) Run() error {
	// Here where you can change your router
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userService := user.NewUserService(server.db)

	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", server.addr)

	return http.ListenAndServe(server.addr, router)
}
