package handlers

import (
	"github.com/71anshuman/go-bookings/pkg/config"
	"github.com/71anshuman/go-bookings/pkg/models"
	"github.com/71anshuman/go-bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Test Hello About!!"

	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
