# NCLOUDS | Docker + Docker Compose

## Database

### Configuration

You need to [install liquibase](https://docs.liquibase.com/install/home.html) before continuing.

Make sure you create the database and run migrations with liquibase before starting the app.

```sql
CREATE DATABASE nclouds_db;
CREATE USER nclouds_user WITH PASSWORD 'nclouds_password';
GRANT ALL PRIVILEGES ON DATABASE nclouds_db to nclouds_user;
```

### Migrations

Inside the db directory use liquibase to run migrations

```sh
cd /backend/db
liquibase --defaultsFile=liquibase.properties --changeLogFile=changelog.xml update
```

Optionally you can validate sql with:

```sh
liquibase --defaultsFile=liquibase.properties --changeLogFile=changelog.xml update
```

## Secrets

Include //pending

`pg_secrets.txt`
```
<<password>>
```