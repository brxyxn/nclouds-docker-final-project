package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brxyxn/go_gpr_nclouds/backend/internal/handlers"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	Router   *mux.Router
	DB       *sql.DB
	L        *log.Logger
	BindAddr string
}

func (a *App) initRoutes() {
	h := handlers.NewHandlers(a.DB)
	// User routes
	a.Router.HandleFunc("/api/v1/users", h.CreateUser).Methods("POST")
	a.Router.HandleFunc("/api/v1/users", h.GetUsers).Methods("GET")
}

/*
To initialize the routes and database connection you must
include the following information as strings and also
call Run setting the port to serve to the web.
*/
func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%v user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	a.DB, err = sql.Open("pgx", connectionStr)
	if err != nil {
		u.Log.Info("Error opening a new connection to the DB.", err)
	}

	a.Router = mux.NewRouter() // Make sure this is set before the server is started

	// We try to validate the connection to the DB is correct, otherwise, the app
	// will restart itself, this is a temporary solution because the postgres image usually
	// is initialized after golang's image.
	err = a.DB.Ping()
	if err != nil {
		a.DB.Close()
		u.Log.Error(err)
	}
	u.Log.Debug("Connected to database", dbname, "with user", user, "at", host+":"+port)
}

/*
Runs the new server.
*/
func (a *App) Run() {
	// Initializing routes
	a.initRoutes()

	// Creating a new server
	srv := http.Server{
		Addr:         a.BindAddr,        // configure the bind address
		Handler:      a.Router,          // set the default handler
		ErrorLog:     a.L,               // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Starting the server
	go func() {
		u.Log.Info("Running server on port", a.BindAddr)

		err := srv.ListenAndServe()
		if err != nil {
			u.Log.Info("Server Status: ", err)
			os.Exit(1)
		}
	}()

	// Creating channel
	cs := make(chan os.Signal, 1)

	signal.Notify(cs, os.Interrupt, os.Kill)
	// signal.Notify(sigchan, os.Kill) // If running on Windows

	sigchan := <-cs
	u.Log.Debug("Signal received:", sigchan)

	ctx, fn := context.WithTimeout(context.Background(), 30*time.Second)
	defer fn()
	srv.Shutdown(ctx)
}
