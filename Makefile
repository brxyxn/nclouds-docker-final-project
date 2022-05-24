# Backend
build:
	go build -o ./backend/bin/main ./backend/cmd/main.go

run_bin:
	./backend/bin/main

env_dev:
	export ENV=Development

env_prod:
	export ENV=Production

backend-build:
	docker build -t go-nclouds-app:latest -f Dockerfile.Backend

backend-run:
	docker run -p 5000:5000 --name nclouds-backend-api --rm -it --env-file .env.local --network hw4_backend go-nclouds-app:latest

# Database
db_validate:
	cd backend/migrations && liquibase --defaults-file=liquibase.properties validate && cd ../..

db_migrate:
	cd backend/migrations && liquibase --defaults-file=liquibase.properties update && cd ../..

psql:
	docker exec -it nclouds-postgres psql -U nclouds_user nclouds_db

# Docker
docker-build:
	docker-compose build

docker-up:
	docker-compose up --build

docker-start:
	docker-compose start

docker-stop:
	docker-compose stop

docker-build-database:
	docker build -t pg-dkf:test -f ./database/Postgres.Dockerfile .

docker-build-cache:
	docker build -t redis-dkf:test -f ./database/Redis.Dockerfile .

docker-exec-psql:
	docker exec -it nclouds-postgres psql -U nclouds_user nclouds_db

docker-exec-bash:
	docker exec -it nclouds-postgres bash

# Frontend
frontend_build:
	docker build -t react-nclouds-app:latest -f Dockerfile.Frontend .

frontend_run:
	docker run -p 3000:80 --name nclouds-frontend-web --rm -it --env-file .env.local react-nclouds-app:latest
