-- CREATE DATABASE nclouds_db;
/**/;
CREATE TABLE IF NOT EXISTS users (
    user_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,

    CONSTRAINT client_pk PRIMARY KEY (user_id)
);