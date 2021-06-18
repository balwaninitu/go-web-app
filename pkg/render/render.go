package render

import (
	"bytes"
	"fmt"
	"go-web-app/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

//go allows to create functions and pass those to templates
var functions = template.FuncMap{}

var app *config.AppConfig

//sets the new config for the template pkg
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, html string) {
	//set configuration so that once have template set. no need to load again
	//untill application restarts
	tc := app.TemplateCache

	//check if template parse, call value out of map

	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	//if template ok, read template and parse, for that create bytes buffer
	//buf will hold information of bytes
	buf := new(bytes.Buffer)

	//put template available in memory in bytes
	//store template in buf variable
	_ = t.Execute(buf, nil)

	//write to response writer, it return byte and err
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browswer", err)
	}

}

//create template map cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	//create map var myCache which will hold the all template
	//mycache will return template with page i.e. home and about which are fully parsed
	//mycache is only inside func scope so add it in return value
	myCache := map[string]*template.Template{}

	//get all template pages that start with anything(*) but end with page

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//for loop for every page which find along with index
	//below page will corresponding to pages variable
	for _, page := range pages {
		//return actual page
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		//create template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//go to templates directory & check if template matches with layout

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		//if there is matche
		if len(matches) > 0 {
			//go to template set and parsed that glob
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		//take that template and add to cache

		myCache[name] = ts

	}
	return myCache, nil
}
