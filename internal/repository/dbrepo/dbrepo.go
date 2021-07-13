package dbrepo

import (
	"database/sql"
	"github.com/71anshuman/go-bookings/internal/config"
	"github.com/71anshuman/go-bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgressRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}