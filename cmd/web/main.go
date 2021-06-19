package main

import (
	"github.com/71anshuman/go-bookings/pkg/config"
	"github.com/71anshuman/go-bookings/pkg/handlers"
	"github.com/71anshuman/go-bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const port = ":9001"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot creat template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Printf("Server is listening on port %s\n", port)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
