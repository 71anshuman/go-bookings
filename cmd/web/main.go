package main

import (
	config2 "github.com/71anshuman/go-bookings/internal/config"
	handlers2 "github.com/71anshuman/go-bookings/internal/handlers"
	render2 "github.com/71anshuman/go-bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const port = ":9001"

var app config2.AppConfig
var session *scs.SessionManager

func main() {
	app.InProd = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	tc, err := render2.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot creat template cache")
	}

	app.TemplateCache = tc

	repo := handlers2.NewRepo(&app)

	handlers2.NewHandler(repo)

	render2.NewTemplate(&app)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Printf("Server is listening on port %s\n", port)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
