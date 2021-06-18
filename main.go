package main

import (
	"fmt"
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"go-web-app/pkg/render"
	"log"

	"net/http"
)

const portNumber = ":8080"

func main() {

	//create variable for config pkg
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	//call func from handlers
	repo := handlers.NewRepo(&app)
	//after crearting repo variable, pass back to handlers
	handlers.NewHandlers(repo)

	//provide access to app config
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting Application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
