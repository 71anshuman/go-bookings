package config

import (
	"html/template"
	"log"

	"github.com/71anshuman/go-bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProd        bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
