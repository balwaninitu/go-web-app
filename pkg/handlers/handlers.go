package handlers

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/models"
	"go-web-app/pkg/render"

	"net/http"
)

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

	/*working with session, grab remote IP address of person visiting
	my site and store it in home page*/
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	//pull remote IP out of session, string will be empty if there is no value
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	//send data to template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
