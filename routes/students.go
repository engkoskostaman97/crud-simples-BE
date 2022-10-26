package routes

import (
	"pelajar/handlers"
	"pelajar/pkg/middleware"
	"pelajar/pkg/mysql"
	"pelajar/repositories"

	"github.com/gorilla/mux"
)

func StudentsRoutes(r *mux.Router) {
	studentRepository := repositories.RepositoryStudent(mysql.DB)
	h := handlers.HandlerStudent(studentRepository)

	r.HandleFunc("/students", h.FindStudents).Methods("GET")
	r.HandleFunc("/students/{id}", h.GetStudent).Methods("GET")
	r.HandleFunc("/students", middleware.Auth(middleware.UploadFile(h.CreateStudent))).Methods("POST") // Write this code
	r.HandleFunc("/students/{id}", middleware.Auth(middleware.UploadFile(h.UpdateStudent))).Methods("PATCH")
	r.HandleFunc("/students/{id}", h.DeleteStudent).Methods("DELETE") // Write this code
}
