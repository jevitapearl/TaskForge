package router

import (
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/handler"
	"github.com/jevitapearl/TaskForge/internal/middleware"
	"github.com/jevitapearl/TaskForge/internal/repository"
)

type Router struct {
	mux *http.ServeMux
}

func New(repo *repository.PostgresRepository) *Router {
	mux := http.NewServeMux()

	h := handler.New(repo)

	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/echo", h.Echo)

	mux.Handle("GET /tasks", middleware.AuthMiddleware(http.HandlerFunc(h.GetAllTasks)))
	mux.Handle("POST /tasks", middleware.AuthMiddleware(http.HandlerFunc(h.CreateTask)))

	mux.Handle("GET /tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.GetTask)))
	mux.Handle("PUT /tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.UpdateTask)))
	mux.Handle("DELETE /tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(h.DeleteTask)))

	mux.HandleFunc("POST /register", h.Register)
	mux.HandleFunc("POST /login", h.Login)
	mux.Handle("POST /refresh", middleware.AuthMiddleware(http.HandlerFunc(h.Refresh)))
	mux.HandleFunc("POST /logout", h.Logout)

	return &Router{
		mux: mux,
	}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
