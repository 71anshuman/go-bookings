package render

import (
	"bytes"
	"fmt"
	config2 "github.com/71anshuman/go-bookings/internal/config"
	models2 "github.com/71anshuman/go-bookings/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config2.AppConfig

func NewTemplate(a *config2.AppConfig) {
	app = a
}

func AddDefaultData(td *models2.TemplateData, r *http.Request) *models2.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplates(w http.ResponseWriter, r *http.Request, tmpl string, td *models2.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./views/*.page.tmpl")

	if err != nil {
		return myCache, nil
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if nil != err {
			return myCache, err
		}

		matches, err := filepath.Glob("./views/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./views/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}
	}
	return myCache, nil
}
