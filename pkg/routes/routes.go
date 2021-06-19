package routes

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
