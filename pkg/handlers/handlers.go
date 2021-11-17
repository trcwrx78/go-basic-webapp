package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/pkg/render"
	"net/http"
)

// The repository used by our handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	 render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."
	// send the data
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
