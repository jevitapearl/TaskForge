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

	mux.Handle("/", middleware.LoggingMiddleware(http.HandlerFunc(h.Home)))
	mux.Handle("/health", middleware.LoggingMiddleware(http.HandlerFunc(h.Health)))
	mux.Handle("/echo", middleware.LoggingMiddleware(http.HandlerFunc(h.Echo)))

	mux.Handle("GET /tasks", middleware.Protected(http.HandlerFunc(h.GetAllTasks)))
	mux.Handle("POST /tasks", middleware.Protected(http.HandlerFunc(h.CreateTask)))

	mux.Handle("GET /tasks/{id}", middleware.Protected(http.HandlerFunc(h.GetTask)))
	mux.Handle("PUT /tasks/{id}", middleware.Protected(http.HandlerFunc(h.UpdateTask)))
	mux.Handle("DELETE /tasks/{id}", middleware.Protected(http.HandlerFunc(h.DeleteTask)))

	mux.Handle("POST /register", middleware.LoggingMiddleware(http.HandlerFunc(h.Register)))
	mux.Handle("POST /login", middleware.LoggingMiddleware(http.HandlerFunc(h.Login)))
	mux.Handle("POST /refresh", middleware.LoggingMiddleware(http.HandlerFunc(h.Refresh)))
	mux.Handle("POST /logout", middleware.LoggingMiddleware(http.HandlerFunc(h.Logout)))

	return &Router{
		mux: mux,
	}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
