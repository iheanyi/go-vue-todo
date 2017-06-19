package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(s *APIService) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos", s.ListTodos).Methods("GET")
	api.HandleFunc("/todos", s.CreateTodo).Methods("POST")
	//s.HandleFunc("/todos/{id}").Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/dist/static"))))
	r.PathPrefix("/").HandlerFunc(IndexHandler())
	return r
}
