package routes

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//recover from panic
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
