package internal

import (
	"database/sql"
	"fmt"

	pg "github.com/brxyxn/go_gpr_nclouds/backend/internal/database/sql/handlers"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func (a *App) sqlRoutes() {
	// SQL Routes
	sqlHandler := pg.NewHandlers(a.DB)
	a.Router.HandleFunc("/api/v1/sql/users", sqlHandler.CreateUser).Methods("POST")
	a.Router.HandleFunc("/api/v1/sql/users", sqlHandler.GetCounter).Methods("GET")
}

/*
To initialize the routes and database connection you must
include the following information as strings and also
call Run setting the port to serve to the web.
*/
func (a *App) InitializePostgresql(host, port, user, password, dbname string) {
	u.Log.Info("Initializing postgres database...")
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
