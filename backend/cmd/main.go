package main

import (
	"fmt"
	"os"
	"strconv"

	a "github.com/brxyxn/go_gpr_nclouds/backend/internal"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func main() {
	a := a.App{}

	a.L = u.InitLogs("nclouds-api ")

	port := os.Getenv("PORT")
	port = fmt.Sprintf("%q", port)
	if port != "" && port != "\"\"" {
		a.BindAddr = ":" + port
	}

	// Env vars
	var db_host, db_port, db_user, db_name, db_password string
	var cache_host, cache_port, cache_password string
	var cache_name int
	var err error
	env := os.Getenv("ENV")
	if env == "Production" {
		// Sql
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USER")
		db_name = os.Getenv("DB_NAME")
		db_password = os.Getenv("DB_PASSWORD")

		// Cache
		cache_host = os.Getenv("RDB_HOST")
		cache_port = os.Getenv("RDB_PORT")
		cache_password = os.Getenv("RDB_PASSWORD")
		cache_name, err = strconv.Atoi(os.Getenv("RDB_NAME"))
		if err != nil {
			u.Log.Error("Error converting env RDB_NAME to int.", err)
		}
	} else {
		// Sql
		db_host = u.DotEnvGet("DB_HOST")
		db_port = u.DotEnvGet("DB_PORT")
		db_user = u.DotEnvGet("DB_USER")
		db_name = u.DotEnvGet("DB_NAME")
		db_password = u.DotEnvGet("DB_PASSWORD")

		// Cache
		cache_host = os.Getenv("RDB_HOST")
		cache_port = os.Getenv("RDB_PORT")
		cache_password = os.Getenv("RDB_PASSWORD")
		cache_name, err = strconv.Atoi(os.Getenv("RDB_NAME"))
		if err != nil {
			u.Log.Error("Error converting env RDB_NAME to int.", err)
		}

		a.BindAddr = ":" + u.DotEnvGet("PORT")
	}

	if db_port == "" || db_host == "" || db_user == "" || db_name == "" || db_password == "" || a.BindAddr == "" {
		u.Log.Error("Environment variables weren't loaded correctly!")
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

	// Cache
	a.InitializeCache(
		cache_host+":"+cache_port,
		cache_password,
		cache_name,
	)

	a.Run()
}
