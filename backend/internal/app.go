package internal

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/cors"
)

type App struct {
	Router   *mux.Router
	DB       *sql.DB
	Cache    *redis.Client
	Ctx      context.Context
	L        *log.Logger
	BindAddr string
	handler  http.Handler
}

func (a *App) initRoutes() {
	a.Router = mux.NewRouter() // Make sure this is set before the server is started
	// SQL Routes
	a.sqlRoutes()

	// Cache Routes
	a.cacheRoutes()

	handler := cors.Default().Handler(a.Router)
	a.handler = handler
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
		Handler:      a.handler,         // set the default handler
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
