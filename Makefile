# Backend
build:
	go build -o ./backend/bin/main ./backend/cmd/main.go

run_bin:
	./backend/bin/main

env_dev:
	export ENV=Development

env_prod:
	export ENV=Production

# Database
db_validate:
	liquibase --defaults-file=./backend/db/liquibase.properties --changelog-file=./backend/db/changelog.xml --classpath=./backend/db/postgresql-42.3.5.jar validate

db_migrate:
	liquibase --defaults-file=./backend/db/liquibase.properties --changelog-file=./backend/db/changelog.xml --classpath=./backend/db/postgresql-42.3.5.jar update

# Docker

# Frontend