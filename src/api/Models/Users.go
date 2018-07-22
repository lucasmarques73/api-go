package models

import "api/Db"

// User struct of table users
type User struct {
	ID     int64  `db:"id,omitempty" json:"id,omitempty"`
	Name   string `validate:"required" db:"name" json:"name,omitempty"`
	Avatar string `db:"avatar" json:"avatar,omitempty"`
	Email  string `validate:"required" db:"email" json:"email,omitempty"`
	Pass   string `validate:"required" db:"pass" json:"-"`
}

// UsersModel Ppinting to the "users" table
var UsersModel = db.Sess.Collection("users")
