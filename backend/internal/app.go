package internal

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
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
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/*")
}

func (a *App) Setup() {
	var err error

	// Env vars
	env := os.Getenv("ENV")
	if env != "Production" {
		err = godotenv.Load()
		if err != nil {
			u.Log.Error("Error loading .env file.", err)
			return
		}
	}

	port := os.Getenv("PORT")
	a.BindAddr = ":" + port

	// Sql
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_sslmode := os.Getenv("DB_SSLMODE")

	// Cache
	cache_host := os.Getenv("RDB_HOST")
	cache_port := os.Getenv("RDB_PORT")
	cache_password := os.Getenv("RDB_PASSWORD")
	cache_name, err := strconv.Atoi(os.Getenv("RDB_NAME"))
	if err != nil {
		u.Log.Error("Error converting env RDB_NAME to int.", err)
	}

	u.Log.Debug("DB Variables:", db_host, db_port, db_user, db_name, db_password)
	u.Log.Debug("Cache Variables:", cache_host, cache_port, cache_name, cache_password)

	if db_port == "" ||
		db_host == "" ||
		db_user == "" ||
		db_name == "" ||
		db_password == "" ||
		a.BindAddr == "" ||
		cache_host == "" ||
		cache_port == "" ||
		cache_name < 0 {
		u.Log.Error("Environment variables weren't loaded correctly!")
		return
	}

	// Sql
	a.initializePostgresql(
		db_host,
		db_port,
		db_user,
		db_password,
		db_name,
		db_sslmode,
	)

	// Cache
	a.initializeCache(
		cache_host+":"+cache_port,
		cache_password,
		cache_name,
	)
}
