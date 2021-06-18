package handlers

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/render"

	"net/http"
)

//holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

//repository used by handlers
var Repo *Repository

/*use repository pattern, it allows swap componants in application at a minimum
changes req to codebase
database connection can also share by using repository in future*/
type Repository struct {
	App *config.AppConfig
}

//creates a new repository
/*it takes pointer to appconfig and populate repository struct with it
return a new instance of it that application holds*/
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//provide access of repository to all handlers by creating method
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	//send data to template
	render.RenderTemplate(w, "about.page.html", &TemplateData{
		StringMap: stringMap,
	})

}
