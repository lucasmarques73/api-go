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
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

// DbSettings app
var dbSettings DbSettings

func GetEnvData(ds DbSettings) DbSettings {

	err := godotenv.Load()
	Errors.CheckErr(err)

	if host := os.Getenv("DB_HOST"); len(host) > 0 {
		ds.DbHost = host
	}
	if database := os.Getenv("DB_DATABASE"); len(database) > 0 {
		ds.DbName = database
	}
	if user := os.Getenv("DB_USER"); len(user) > 0 {
		ds.DbUser = user
	}
	if password := os.Getenv("DB_PASSWORD"); len(password) > 0 {
		ds.DbPassword = password
	}
	if port := os.Getenv("DB_PORT"); len(port) > 0 {
		ds.DbPort = port
	}
	return ds
}

// Sess Database connection
var Sess db.Database

func init() {
	var err error

	dbSettings = GetEnvData(dbSettings)

	var settings = postgresql.ConnectionURL{
		User:     dbSettings.DbUser,
		Password: dbSettings.DbPassword,
		Host:     dbSettings.DbHost,
		Database: dbSettings.DbName,
	}

	Sess, err = postgresql.Open(settings)
	Errors.CheckErr(err)
	Sess.SetLogging(true)

	fmt.Println("Database successfully connected")
}
