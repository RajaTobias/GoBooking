package handlers

import (
	"net/http"

	"github.com/RajaTobias/GoBookings/pkg/config"
	"github.com/RajaTobias/GoBookings/pkg/models"
	"github.com/RajaTobias/GoBookings/pkg/render"
)

// hold the repository
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// make a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// make a new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	stringMap["test"] = "Find more about this page"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
