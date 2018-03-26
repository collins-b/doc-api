package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App structure
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize db connection
func (a *App) Initialize(user, password, dbName string) {}

// Run app
func (a *App) Run(addr string) {}
