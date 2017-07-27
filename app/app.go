package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"./models"
)

// Main class implementing whole application.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) InitDB(user, password, dbname string) {
	models.InitDB(
		"postgres",
		"postgres",
		"postgres",
	)
}

// Method start app.
func (a *App) Run(addr string) {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(addr, router))
}
