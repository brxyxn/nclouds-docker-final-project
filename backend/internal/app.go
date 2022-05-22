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

	rs "github.com/brxyxn/go_gpr_nclouds/backend/internal/database/cache/handlers"
	pg "github.com/brxyxn/go_gpr_nclouds/backend/internal/database/sql/handlers"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	Router   *mux.Router
	DB       *sql.DB
	Cache    *redis.Client
	Ctx      context.Context
	L        *log.Logger
	BindAddr string
}

func (a *App) initRoutes() {
	a.Router = mux.NewRouter() // Make sure this is set before the server is started
	// SQL Routes
	sqlHandler := pg.NewHandlers(a.DB)
	a.Router.HandleFunc("/api/v1/sql/users", sqlHandler.CreateUser).Methods("POST")

	// Cache Routes
	cacheHandler := rs.NewHandlers(a.Cache)
	a.Router.HandleFunc("/api/v1/cache/users", cacheHandler.CreateUser).Methods("POST")
}

/*
To initialize the routes and database connection you must
include the following information as strings and also
call Run setting the port to serve to the web.
*/
func (a *App) InitializePostgresql(host, port, user, password, dbname string) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%v user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	a.DB, err = sql.Open("pgx", connectionStr)
	if err != nil {
		u.Log.Error("Error opening a new connection to the DB.", err)
		return
	}

	// We try to validate the connection to the DB is correct, otherwise, the app
	// will restart itself, this is a temporary solution because the postgres image usually
	// is initialized after golang's image.
	err = a.DB.Ping()
	if err != nil {
		a.DB.Close()
		u.Log.Error(err)
		return
	}
	u.Log.Info("Connected to database", dbname, "with user", user, "at", host+":"+port)
}

func (a *App) InitializeCache(bindAddr, password string, dbname int) {
	a.Ctx = context.Background()
	a.Cache = redis.NewClient(&redis.Options{
		Addr:     bindAddr,
		Password: password,
		DB:       dbname,
	})

	ping := a.Cache.Ping(a.Ctx)
	if ping.Val() != "PONG" {
		u.Log.Error("Error opening a new connection to the Redis Cache.")
		return
	}
	u.Log.Info("Connected to redis cache db", dbname, "at", bindAddr)
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
