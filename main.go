package main

import (
	"fmt"
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"go-web-app/pkg/render"
	"go-web-app/pkg/routes"
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

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(&app),
	}
	fmt.Printf("Starting Application on port %s\n", portNumber)
	//start actual server
	err = srv.ListenAndServe()
	log.Fatal(err)

}
