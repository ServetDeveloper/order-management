package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ServetDeveloper/order-management/service/user"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)

	userHandler.RegisterRoutes(subRouter)

	log.Printf("Listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
