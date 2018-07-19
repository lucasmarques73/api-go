CREATE DATABASE go;

\c go

CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    pass VARCHAR(255) NOT NULL
)

CREATE TABLE widgets (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(50),
    price NUMERIC,
    melts BOOLEAN,
    inventory INT
)