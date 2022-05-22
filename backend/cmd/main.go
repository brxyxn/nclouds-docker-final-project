package main

import (
	"fmt"
	"os"

	b "github.com/brxyxn/go_gpr_nclouds/backend/internal"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func main() {
	a := b.App{}

	a.L = u.InitLogs("nclouds-api ")

	port := os.Getenv("PORT")
	port = fmt.Sprintf("%q", port)
	if port != "" && port != "\"\"" {
		a.BindAddr = ":" + port
	}

	var db_host, db_port, db_user, db_name, db_password string = "", "", "", "", ""
	env := os.Getenv("ENV")
	if env == "Production" {
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USER")
		db_name = os.Getenv("DB_NAME")
		db_password = os.Getenv("DB_PASSWORD")
	} else {
		db_host = u.DotEnvGet("DB_HOST")
		db_port = u.DotEnvGet("DB_PORT")
		db_user = u.DotEnvGet("DB_USER")
		db_name = u.DotEnvGet("DB_NAME")
		db_password = u.DotEnvGet("DB_PASSWORD")

		a.BindAddr = ":" + u.DotEnvGet("PORT")
	}

	if db_port == "" || db_host == "" || db_user == "" || db_name == "" || db_password == "" || a.BindAddr == "" {
		u.Log.Error("Environment variables were not loaded correctly!")
		return
	}

	// Sql
	a.InitializePostgresql(
		db_host,
		db_port,
		db_user,
		db_password,
		db_name,
	)

	a.Run()
}
