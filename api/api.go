package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iheanyi/go-vue-todo/db"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
type CreateTodoRequest struct {
	Description string `json:"description"`
}

type ListTodosResponse struct {
	Todos []*db.Todo `json:"todos"`
}

type APIService struct {
	db db.Database
}

func (s *APIService) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createTodo CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&createTodo); err != nil {
		panic(err)
	}
	todo := &db.Todo{
		Description: createTodo.Description,
	}
	defer r.Body.Close()
	todo, err := s.db.CreateTodo(todo)
	if err != nil {
		renderError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, todo, http.StatusOK)
}

func (s *APIService) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.db.ListTodos()
	if err != nil {
		// Return error here
		log.Printf("error listing todos: %v", err)
		renderError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todosResponse := &ListTodosResponse{
		Todos: todos,
	}

	renderJSON(w, todosResponse, http.StatusOK)
	return
}

func NewAPIService(db db.Database) *APIService {
	return &APIService{
		db: db,
	}
}

func IndexHandler() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/dist/index.html")
	}

	return http.HandlerFunc(fn)
}

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

func renderJSON(w http.ResponseWriter, res interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("error encoding json: %v", err)
		panic(err)
	}
}

func renderError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := &ErrorResponse{
		Error: err,
		Code:  code,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("error encoding json: %v", err)
		panic(err)
	}
}
