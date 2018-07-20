package main

import (
	"api/errors"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var createTableUsers = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    pass VARCHAR(255) NOT NULL
)`

var createTableWidgets = `CREATE TABLE IF NOT EXISTS widgets (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(50),
    price NUMERIC,
    melts BOOLEAN,
    inventory INT
)`

func main() {
	var conStr = "host=db port=5432 user=go password=go dbname=go sslmode=disable"
	var db, err = sql.Open("postgres", conStr)
	errors.CheckErr(err)

	exec(db, createTableUsers)
	exec(db, createTableWidgets)
	db.Close()
	fmt.Println("Successfully created database")
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	errors.CheckErr(err)
	return result
}
