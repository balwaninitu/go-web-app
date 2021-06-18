package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

//go allows to create functions and pass those to templates
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, html string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting in template cache", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("cannot parsed templates", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
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
