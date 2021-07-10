package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/71anshuman/go-bookings/internal/config"
	"github.com/71anshuman/go-bookings/internal/handlers"
	"github.com/71anshuman/go-bookings/internal/models"
	"github.com/71anshuman/go-bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":9001"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Printf("Server is listening on port %s\n", port)

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

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
		return err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProd

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplate(&app)

	return nil
}
