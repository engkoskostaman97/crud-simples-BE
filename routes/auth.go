package routes

import (
	"pelajar/handlers"
	"pelajar/pkg/mysql"
	"pelajar/repositories"

	"pelajar/pkg/middleware"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/checkauth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
