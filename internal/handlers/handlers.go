package handlers

import (
	"encoding/json"
	"fmt"
	config2 "github.com/71anshuman/go-bookings/internal/config"
	models2 "github.com/71anshuman/go-bookings/internal/models"
	render2 "github.com/71anshuman/go-bookings/internal/render"
	"log"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config2.AppConfig
}

func NewRepo(a *config2.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplates(w,r,  "home.page.tmpl",&models2.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplates(w, r, "make-reservation.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplates(w,r,  "generals.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplates(w,r, "majors.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplates(w,r, "search-availability.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil{
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render2.RenderTemplates(w,r, "contact.page.tmpl", &models2.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Test Hello About!!"

	render2.RenderTemplates(w,r, "about.page.tmpl", &models2.TemplateData{
		StringMap: stringMap,
	})
}
