package Db

import (
	"api/Errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	db "upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

// DbSettings app
type DbSettings struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
}

// DbSettings app
var dbSettings DbSettings

func GetEnvData(dbSettings DbSettings) DbSettings {

	err := godotenv.Load()
	Errors.CheckErr(err)

	if host := os.Getenv("DB_HOST"); len(host) > 0 {
		dbSettings.dbHost = host
	}
	if database := os.Getenv("DB_DATABASE"); len(database) > 0 {
		dbSettings.dbName = database
	}
	if user := os.Getenv("DB_USER"); len(user) > 0 {
		dbSettings.dbUser = user
	}
	if password := os.Getenv("DB_PASSWORD"); len(password) > 0 {
		dbSettings.dbPassword = password
	}
	if port := os.Getenv("DB_PORT"); len(port) > 0 {
		dbSettings.dbPort = port
	}
	return dbSettings
}

// Sess Database connection
var Sess db.Database

func init() {
	var err error

	dbSettings = GetEnvData(dbSettings)

	var settings = postgresql.ConnectionURL{
		User:     dbSettings.dbUser,
		Password: dbSettings.dbPassword,
		Host:     dbSettings.dbHost,
		Database: dbSettings.dbName,
	}

	Sess, err = postgresql.Open(settings)
	Errors.CheckErr(err)
	Sess.SetLogging(true)

	fmt.Println("Database successfully connected")
}
