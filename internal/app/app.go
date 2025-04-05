package app

import (
	"habit-tracker-api/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func NewApp() *App {
	r := mux.NewRouter()
	handlers.RegisterListRoutes(r)
	handlers.RegisterTaskRoutes(r)
	return &App{router: r}
}

func (a *App) Run(addr string) error {
	return http.ListenAndServe(addr, a.router)
}
