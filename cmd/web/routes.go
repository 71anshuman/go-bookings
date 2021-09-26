package main

import (
	"net/http"

	"github.com/71anshuman/go-bookings/internal/config"
	"github.com/71anshuman/go-bookings/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/logout", handlers.Repo.GetLogout)
	mux.Get("/login", handlers.Repo.GetLogin)
	mux.Post("/login", handlers.Repo.PostLogin)

	mux.Post("/clear", handlers.Repo.PostClear)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservations", handlers.Repo.Reservation)
	mux.Post("/make-reservations", handlers.Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)

		mux.Get("/dashboard", handlers.Repo.AdminDashboard)
	})

	return mux
}
