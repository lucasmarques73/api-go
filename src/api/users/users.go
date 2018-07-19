package users

import (
	"database/sql"
	_ "github.com/lib/pq"
	"api/errors"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	Pass string `json:"-"`
}

var conStr = "host=db port=5432 user=go password=go dbname=go sslmode=disable"
var db,err = sql.Open("postgres", conStr)

func ListUsers() []User {

	rows, err := db.Query("SELECT * FROM users")
	errors.CheckErr(err)

	users := []User{}

	for rows.Next(){
		u := User{}
		rows.Scan(&u.Id, &u.Name, &u.Avatar,&u.Email,&u.Pass)
		users = append(users,u)
	}

	return users
}

func GetUserById(id string) User {
	
	row:= db.QueryRow("SELECT * FROM users WHERE id=?",id)
	u := User{}
	row.Scan(&u.Id, &u.Name, &u.Avatar,&u.Email,&u.Pass)
	
	return u
}

func CreateUser(user User) User {
	return user
}

func UpdateUser(user User, id string) User {
	return user
}