package users

// import "errors"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	Pass string `json:"-"`
}

func ListUsers() []User {

	users := []User {
		User{Id:1,Name:"teste 1",Avatar:"teste",Email:"teste1@hotmail.com",Pass:"123"},
		User{Id:2,Name:"teste 2",Avatar:"teste",Email:"teste2@hotmail.com",Pass:"123"},
		User{Id:3,Name:"teste 3",Avatar:"teste",Email:"lucas3@hotmail.com",Pass:"123"},
	}

	return users
}

func GetUserById(id string) User {
	user := User{Id:1,Name:"teste 1",Avatar:"teste",Email:"teste1@hotmail.com",Pass:"123"}
	return user
}

func CreateUser(user User) User {
	user.Id = 4
	user.Name = "Att"
	return user
}

func UpdateUser(user User, id string) User {
	return user
}