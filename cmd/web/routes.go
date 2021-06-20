package main

import (
	config2 "github.com/71anshuman/go-bookings/internal/config"
	handlers2 "github.com/71anshuman/go-bookings/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config2.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers2.Repo.Home)
	mux.Get("/about", handlers2.Repo.About)
	mux.Get("/generals-quarters", handlers2.Repo.Generals)
	mux.Get("/majors-suite", handlers2.Repo.Majors)

	mux.Get("/search-availability", handlers2.Repo.Availability)
	mux.Post("/search-availability", handlers2.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers2.Repo.AvailabilityJSON)

	mux.Get("/contact", handlers2.Repo.Contact)

	mux.Get("/make-reservations", handlers2.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	return mux
}
