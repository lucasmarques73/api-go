package users

import (
	"api/errors"
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

type User struct {
	ID     int    `db:"id" json:"id,omitempty"`
	Name   string `db:"name" json:"name,omitempty"`
	Avatar string `db:"avatar" json:"avatar,omitempty"`
	Email  string `db:"email" json:"email,omitempty"`
	Pass   string `db:"pass" json:"-"`
}

var conStr = "host=db port=5432 user=go password=go dbname=go sslmode=disable"
var db, err = sql.Open("postgres", conStr)

func ListUsers() []User {

	rows, err := db.Query("SELECT * FROM users")
	errors.CheckErr(err)

	users := []User{}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.Avatar, &user.Email, &user.Pass)
		users = append(users, user)
	}

	return users
}

func GetUserById(idS string) User {
	id, _ := strconv.Atoi(idS)
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	user := User{}
	row.Scan(&user.ID, &user.Name, &user.Avatar, &user.Email, &user.Pass)

	return user
}

func CreateUser(user User) User {
	return user
}

func UpdateUser(user User, idS string) User {
	return user
}

func DeleteUser(id string) {

}
