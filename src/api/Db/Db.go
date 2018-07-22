package Db

import (
	"api/Errors"
	"fmt"

	"upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

var settings = postgresql.ConnectionURL{
	User:     "go",
	Password: "go",
	Host:     "db",
	Database: "go",
}

// Sess Database connection
var Sess db.Database

func init() {
	var err error

	Sess, err = postgresql.Open(settings)
	Errors.CheckErr(err)
	Sess.SetLogging(true)

	fmt.Println("Database successfully connected")
}
