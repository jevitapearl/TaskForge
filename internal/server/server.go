package server

import (
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/repository"
	"github.com/jevitapearl/TaskForge/internal/service"
)

type Server struct {
	mux     *http.ServeMux
	service *service.TaskService
}

func New() *Server {
	s := &Server{
		mux:     http.NewServeMux(),
		service: service.New(repository.New()),
	}

	s.routes()

	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/", s.Home)
	s.mux.HandleFunc("/health", s.Health)
	s.mux.HandleFunc("/echo", s.Echo)

	s.mux.HandleFunc("GET /tasks/{id}", s.GetTask)
	s.mux.HandleFunc("PUT /tasks/{id}", s.UpdateTask)
	s.mux.HandleFunc("DELETE /tasks/{id}", s.DeleteTask)

	s.mux.HandleFunc("GET /tasks", s.GetAllTasks)
	s.mux.HandleFunc("POST /tasks", s.CreateTask)
}

func (s *Server) Router() http.Handler {
	return s.mux
}
