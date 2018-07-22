package Controllers

import (
	"api/Errors"
	"api/Models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllUsers -  Listing all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var users []Models.User
	if err := Models.UsersModel.Find().All(&users); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "Users not found",
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Errors:  false,
			Data:    users,
			Message: "List of all users",
		})
	}
}

// GetUser - Listing a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	idS := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idS)

	var user Models.User
	res := Models.UsersModel.Find(id)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "User not found",
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Errors:  false,
			Data:    user,
			Message: "User data of id " + idS,
		})
	}
}

// CreateUser - Creating a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var user Models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	e := Models.UsersModel.Find("email", user.Email)

	if count, _ := e.Count(); count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "The email field must be unique",
		})
	} else {
		if res, err := Models.UsersModel.Insert(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Errors.CheckErr(err)
		} else {
			user.ID = res.(int64)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Response{
				Errors:  false,
				Data:    user,
				Message: "User created",
			})
		}
	}

}

// UpdateUser - Updating a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var user Models.User
	res := Models.UsersModel.Find(id)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  true,
			Data:    "",
			Message: "User not found",
		})
	} else {
		_ = json.NewDecoder(r.Body).Decode(&user)

		// Validate duplicate email
		// eid Email and ID
		e := Models.UsersModel.Find("email", user.Email)
		count, _ := e.Count()
		var eid Models.User
		e.One(&eid)

		if count > 0 && user.ID != eid.ID {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Errors:  true,
				Data:    "",
				Message: "The email field must be unique",
			})
		} else {
			if err = res.Update(user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Errors.CheckErr(err)
			} else {
				json.NewEncoder(w).Encode(Response{
					Errors:  false,
					Data:    user,
					Message: "User updated",
				})
			}
		}
	}
}

// DeleteUser - Deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var user Models.User
	res := Models.UsersModel.Find(id)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Errors:  false,
			Data:    "",
			Message: "User not found",
		})
	} else {
		_ = json.NewDecoder(r.Body).Decode(&user)

		if err = res.Delete(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Errors.CheckErr(err)
		} else {
			json.NewEncoder(w).Encode(Response{
				Errors:  false,
				Data:    user,
				Message: "User deleted",
			})
		}
	}
}
