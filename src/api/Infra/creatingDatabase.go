package main

import (
	"api/errors"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/icrowley/fake"
	_ "github.com/lib/pq"
)

var createTableUsersSQL = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    pass VARCHAR(255) NOT NULL
)`

var createTableWidgetsSQL = `CREATE TABLE IF NOT EXISTS widgets (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    color VARCHAR(50),
    price NUMERIC,
    melts BOOLEAN,
    inventory INT NOT NULL
)`

var usersTableSeedSQL = `INSERT INTO users (name,avatar,email,pass) VALUES (?,?,?,?)`
var widgetsTableSeedSQL = `INSERT INTO widgets (name,color,price,melts,inventory) VALUES (?,?,?,?,?)`

func main() {
	var conStr = "host=db port=5432 user=go password=go dbname=go sslmode=disable"
	var db, err = sql.Open("postgres", conStr)
	errors.CheckErr(err)

	createTableUsers(db)
	createTableWidgets(db)
	usersTableSeed(db)
	widgetsTableSeed(db)

	db.Close()
	fmt.Println("Successfully created database")
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	errors.CheckErr(err)
	return result
}

func createTableUsers(db *sql.DB) {
	exec(db, createTableUsersSQL)
}

func createTableWidgets(db *sql.DB) {
	exec(db, createTableWidgetsSQL)
}

func usersTableSeed(db *sql.DB) {

	for index := 0; index < 5; index++ {
		name := fake.FullName()
		avatar := "https://loremflickr.com/320/240/cats"
		email := fake.EmailAddress()
		pass := fake.SimplePassword()

		exec(db, "INSERT INTO users (name,avatar,email,pass) VALUES ('"+name+"','"+avatar+"','"+email+"','"+pass+"')")

	}
}

func widgetsTableSeed(db *sql.DB) {

	for index := 0; index < 5; index++ {
		name := fake.FullName()
		color := fake.Color()
		price := strconv.FormatFloat(randomdata.Decimal(100), 'E', 2, 64)
		melts := strconv.FormatBool(randomdata.Boolean())
		inventory := strconv.Itoa(randomdata.Number(50))

		exec(db, "INSERT INTO widgets (name,color,price,melts,inventory) VALUES ('"+name+"','"+color+"','"+price+"','"+melts+"','"+inventory+"')")

	}
}
