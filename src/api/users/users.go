package users

import (
	"database/sql"
	_ "github.com/lib/pq"
	"api/errors"
)

type User struct {
	Id int `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name,omitempty"`
	Avatar string `db:"avatar" json:"avatar,omitempty"`
	Email string `db:"email" json:"email,omitempty"`
	Pass string `db:"pass" json:"-"`
}

var conStr = "host=db port=5432 user=go password=go dbname=go sslmode=disable"
var db,err = sql.Open("postgres", conStr)

func ListUsers() []User {

	rows, err := db.Query("SELECT * FROM users")
	errors.CheckErr(err)

	users := []User{}

	for rows.Next(){
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Avatar,&user.Email,&user.Pass)
		users = append(users,user)
	}

	return users
}

func GetUserById(id string) User {
	
	row:= db.QueryRow("SELECT * FROM users WHERE id=?",id)
	user := User{}
	row.Scan(&user.Id, &user.Name, &user.Avatar,&user.Email,&user.Pass)
	
	return user
}

func CreateUser(user User) User {
	return user
}

func UpdateUser(user User, id string) User {
	return user
}