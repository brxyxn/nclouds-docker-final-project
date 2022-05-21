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
move_fwd:
	cd backend/db
move_bwd:
	cd ../..

db_validate:
	cd backend/db && liquibase --defaults-file=liquibase.properties validate && cd ../..

db_migrate:
	cd backend/db && liquibase --defaults-file=liquibase.properties update && cd ../..

psql:
	docker exec -it nclouds-postgres psql -U nclouds_user nclouds_db

# Docker
docker-build:
	docker-compose build

docker-start:
	docker-compose start

docker-stop:
	docker-compose stop

docker-build-database:
	docker build -t pg-dkf:test -f ./database/Postgres.Dockerfile .

docker-build-cache:
	docker build -t redis-dkf:test -f ./database/Redis.Dockerfile .

# Frontend