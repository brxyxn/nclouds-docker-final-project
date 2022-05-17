package main

import (
	"flag"
	"log"
	"os"

	b "github.com/brxyxn/go_gpr_nclouds/backend"
	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

func main() {
	flag.Parse()

	a := b.App{}

	u.InitLogs("go-gpr-api ")
	a.L = log.New(os.Stdout, "go-gpr-api ", log.LstdFlags)

	port := os.Getenv("PORT")
	if port == "" {
		a.BindAddr = ":" + "3000"
	} else {
		a.BindAddr = ":" + port
	}
	u.LogInfo(a.BindAddr)

	var db_host, db_port, db_user, db_password, db_name = "", "", "", "", ""
	env := os.Getenv("ENV")
	if env == "Production" {
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USER")
		db_password = os.Getenv("DB_PASSWORD")
		db_name = os.Getenv("DB_NAME")
	} else if env == "Development" {
		db_host = u.DotEnvGet("PG_HOST")
		db_port = os.Getenv("PG_PORT")
		db_user = os.Getenv("PG_USER")
		db_password = os.Getenv("PG_PASSWORD")
		db_name = os.Getenv("PG_NAME")
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
