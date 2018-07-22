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

// var ConStr = "host=db port=5432 widget=go password=go dbname=go sslmode=disable"

// func ListUsers() ([]User, error) {
// 	var db, _ = sql.Open("postgres", ConStr)
// 	rows, err := db.Query("SELECT id,name,avatar,email FROM users")
// 	// errors.CheckErr(err)

// 	users := []User{}

// 	for rows.Next() {
// 		user := User{}
// 		err = rows.Scan(&user.ID, &user.Name, &user.Avatar, &user.Email)
// 		users = append(users, user)
// 	}

// 	return users, err
// }

// func GetUserById(id string) (User, error) {
// 	var db, _ = sql.Open("postgres", ConStr)
// 	user := User{}
// 	err := db.QueryRow("SELECT id,name,avatar,email FROM users WHERE id ="+id).Scan(&user.ID, &user.Name, &user.Avatar, &user.Email)
// 	db.Close()

// 	return user, err
// }

// func CreateUser(user User) (User, error) {
// 	var db, _ = sql.Open("postgres", ConStr)
// 	err := db.QueryRow("INSERT INTO users (name,avatar,email,pass) VALUES ('" + user.Name + "','" + user.Avatar + "','" + user.Email + "','" + user.Pass + "') RETURNING id").Scan(&user.ID)
// 	db.Close()

// 	return user, err
// }

// func UpdateUser(user User) (User, error) {
// 	var db, _ = sql.Open("postgres", ConStr)
// 	_, err := db.Exec("UPDATE users SET name=" + user.Name + ",avatar=" + user.Avatar + ",email=" + user.Email + " WHERE id=" + strconv.FormatInt(user.ID, 10))
// 	db.Close()

// 	return user, err
// }

// func DeleteUser(user User) error {
// 	var db, _ = sql.Open("postgres", ConStr)
// 	_, err := db.Exec("DELETE FROM users WHERE id = " + strconv.FormatInt(user.ID, 10))
// 	db.Close()

// 	return err
// }
