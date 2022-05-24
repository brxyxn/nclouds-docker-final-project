# nClouds | Docker | Final Project

This project is the final project for Docker class taught by [nClouds Academy](https://www.nclouds.com/).

Tech Stack used in this project:

- [Golang 1.18](https://go.dev)
- [PostgreSQL 14.2](https://www.postgresql.org)
- [Docker 20.10.16](https://www.docker.com/get-started)
- [Docker Compose 2.5.0](https://docs.docker.com/compose/)
- [Ubuntu](https://ubuntu.com/)

> _Note: `make` commands should be executed from project's root directory_ > _Take in consideration that this project is intended to use two layers the backend api and frontend as an independent part in order to manage the server side load and performance_

Links to Docker Hub

- [React App](https://hub.docker.com/r/brxyxn/react-nclouds-app)
- [Golang App](https://hub.docker.com/r/brxyxn/go-nclouds-app)

## Preview

![preview](./preview/cover.png)

# Using Docker Compose

## Clone

### GitHub Repository

```sh
git clone git@github.com:brxyxn/nclouds-docker-final-project.git
cd nclouds-docker-final-project
```

### Docker Hub Repository

```sh
docker pull brxyxn/react-nclouds-app
docker pull brxyxn/go-nclouds-app
```

## Docker Swarm

Initialize docker swarm

```sh
docker swarm init
```

In case you have multiple network ids, please specify the `--advertise-addr` flag.

```sh
docker swarm init --advertise-addr # format: <ip|interface>[:port]
```

## Docker Secrets

Create your secrets file and specify a secret password, see example:

```sh
echo "secret_password" > pg_secrets.txt
```

`pg_secrets.txt`

```txt
secret_password
```

## Execute

```sh
docker-compose up --build
# or
docker compose up --build
# or
make docker-up
```

If you want to start the containers with no output first do `build` and then `start`

```sh
docker-compose build
docker-compose start
# or
make docker-build
make docker-start
```

## Open

Open your browser and enter to [localhost:3000](http://localhost:3000/)

### Screenshots

> Password Component Show Password Off

![Password Component Show Off](./preview/password-off.png)

> Password Component Show Password On

![Password Component Show On](./preview/password-on.png)

> Item Saved to Cache

![Updated items](./preview/update.png)

# Using Docker Images

## Details

The first step you need to know is that you can configure your own database container or existing installation in your local machine, either option you can follow this steps.

## Env file

Create the .env file to pass the environment variables to the docker images.

Remember to update the values with its respective values, if you are running `postgres` and `redis` with docker, please consider using a specific network with the `--network` flag and then use `docker inspect [container-name] | grep IPAddress` and use it as `_HOST`. You also need to configure `user, password and db_name` and additionaly create the **users table** you can [see schema](#configuration) below.

`./.env`

```.env
# Backend
## Golang
PORT=5000

### Postgres
DB_HOST=172.21.0.3
DB_PORT=5432
DB_USER=nclouds_user
DB_PASSWORD=secret
DB_NAME=nclouds_db
DB_DRIVER=pgx
# required
DB_SSLMODE=disable

### Redis
RDB_HOST=172.21.0.2
RDB_PORT=6379
RDB_PASSWORD=
RDB_NAME=0

REACT_APP_API_URL=http://localhost:5000/api/v1
```

## Pull the images

```sh
docker pull brxyxn/react-nclouds-app
docker pull brxyxn/go-nclouds-app
```

Or use the `Dockerfile.Backend` and `Dockerfile.Frontend` files and build the images locally

`for backend`

```sh
docker build -t go-nclouds-app:latest -f Dockerfile.Backend
docker run -p 5000:5000 --name nclouds-backend-api --rm -it --env-file .env --network hw4_backend brxyxn/go-nclouds-app:latest
```

> Please note the `--network` flag is being used to make sure the container will be able to **communicate** with the `database` and `cache` engines, unless you are using your local installation of `postgres` and `redis` you don't need it.

`for frontend`

```
docker build -t react-nclouds-app:latest -f Dockerfile.Frontend .
docker run -p 3000:80 --name nclouds-frontend-web --rm -it --env-file .env brxyxn/react-nclouds-app:latest
```

_If you cloned the repository you can make it with make._

```sh
# backend first
make backend-build
make backend-run
# frontend last
make frontend_build
make frontend_run
```

## Database

### Configuration

Make sure the database is created when running `docker-compose up` the first time by reading the output. By default the file is included and mounted to the database and it should be setup automatically, if any error ocurrs you can create the database with the following script.

`database/setup.sql`

```sql
-- CREATE DATABASE nclouds_db;
CREATE TABLE IF NOT EXISTS users (
    user_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,

    CONSTRAINT client_pk PRIMARY KEY (user_id)
);
```

### Migrations (Optional)

You need to [install liquibase](https://docs.liquibase.com/install/home.html) before continuing.

Update `username` and `password` values in file `backend/migrations/liquibase.properties` based on `DC_PG_USER` and `pg_secret.txt` values.

```properties
# ...
username: nclouds_user
password: secret
# ...
```

Inside the `migrations` directory use liquibase to run migrations

```sh
cd /backend/migrations
liquibase --defaultsFile=liquibase.properties --changeLogFile=changelog.xml update
# or
make db_migrate
```

Optionally you can validate migrations with `validate`

```sh
cd /backend/migrations
liquibase --defaultsFile=liquibase.properties --changeLogFile=changelog.xml validate
# or
make db_validate
```
