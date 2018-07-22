package models

import "api/db"

// User struct of table users
type User struct {
	ID     int64  `db:"id" json:"id,omitempty"`
	Name   string `db:"name" json:"name,omitempty"`
	Avatar string `db:"avatar" json:"avatar,omitempty"`
	Email  string `db:"email" json:"email,omitempty"`
	Pass   string `db:"pass" json:"pass,omitempty"`
}

// UsersModel Ppinting to the "users" table
var UsersModel = db.Sess.Collection("users")
