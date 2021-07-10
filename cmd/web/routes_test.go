package main

import (
	"fmt"
	"testing"

	"github.com/71anshuman/go-bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// Test pass do nothing
	default:
		t.Error(fmt.Sprintf("Type is %T, expected type *chi.Mux", v))
	}
}
