package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/71anshuman/go-bookings/internal/driver"
	"github.com/71anshuman/go-bookings/internal/helpers"

	"github.com/71anshuman/go-bookings/internal/config"
	"github.com/71anshuman/go-bookings/internal/handlers"
	"github.com/71anshuman/go-bookings/internal/models"
	"github.com/71anshuman/go-bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":9001"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Printf("Server is listening on port %s\n", port)

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	app.InProd = false

	infoLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	app.Session = session

	// connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 user=anshumanlawania password=")
	if err != nil {
		log.Fatal("Cannot connect to DB!!")
	}
	log.Println("Connected to DB")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot creat template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProd

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
