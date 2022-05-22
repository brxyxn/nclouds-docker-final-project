package main

import (
	"fmt"
	"log"
	"os"

	b "github.com/brxyxn/go_gpr_nclouds/backend"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func main() {
	a := b.App{}

	u.InitLogs("go-gpr-api ")
	a.L = log.New(os.Stdout, "go-gpr-api ", log.LstdFlags)

	port := os.Getenv("PORT")
	port = fmt.Sprintf("%q", port)
	if port == "" || port == "\"\"" {
		a.BindAddr = ":" + "3000"
	} else {
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
	} else if env == "Development" {
		db_host = u.DotEnvGet("DB_HOST")
		db_port = u.DotEnvGet("DB_PORT")
		db_user = u.DotEnvGet("DB_USER")
		db_name = u.DotEnvGet("DB_NAME")
		db_password = u.DotEnvGet("DB_PASSWORD")
	}

	a.Initialize(
		db_host,
		db_port,
		db_user,
		db_password,
		db_name,
	)

	a.Run()
}
