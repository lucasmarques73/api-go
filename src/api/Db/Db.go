package Db

import (
	"api/Errors"
	"api/Services/GetEnvData"
	"fmt"

	db "upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

// Sess Database connection
var Sess db.Database

func init() {
	var err error
	var dbSettings GetEnvData.Settings
	dbSettings = GetEnvData.GetEnvData(dbSettings)

	var settings = postgresql.ConnectionURL{
		User:     dbSettings.DbUser,
		Password: dbSettings.DbPassword,
		Host:     dbSettings.DbHost,
		Database: dbSettings.DbName,
	}

	Sess, err = postgresql.Open(settings)
	Errors.CheckErr(err)
	Sess.SetLogging(dbSettings.DbLogging)

	fmt.Println("Database successfully connected")
}
