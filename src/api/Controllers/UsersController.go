package controllers

import (
	"api/Errors"
	"api/Models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var users []models.User
	if err := models.UsersModel.Find().All(&users); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"errors":  strconv.FormatBool(true),
			"data":    "",
			"message": "Users not found",
		})
	} else {
		data, _ := json.Marshal(users)
		json.NewEncoder(w).Encode(map[string]string{
			"errors":  strconv.FormatBool(false),
			"data":    string(data),
			"message": "List of all users",
		})
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	idS := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idS)

	var user models.User
	res := models.UsersModel.Find(id)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(true),
			"data":    "",
			"message": "User not found",
		})
	} else {
		data, _ := json.Marshal(user)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(false),
			"data":    string(data),
			"message": "User data of id " + idS,
		})
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	e := models.UsersModel.Find("email", user.Email)

	if count, _ := e.Count(); count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(true),
			"data":    "",
			"message": "The email fild must be unique",
		})
	} else {
		if res, err := models.UsersModel.Insert(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errors.CheckErr(err)
		} else {
			user.ID = res.(int64)
			w.WriteHeader(http.StatusCreated)
			data, _ := json.Marshal(user)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   strconv.FormatBool(false),
				"data":    string(data),
				"message": "User created",
			})
		}
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var user models.User
	res := models.UsersModel.Find(id)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(true),
			"data":    "",
			"message": "User not found",
		})
	} else {
		_ = json.NewDecoder(r.Body).Decode(&user)

		// Validate duplicate email
		e := models.UsersModel.Find("email", user.Email)
		count, _ := e.Count()
		var eid models.User
		e.One(&eid)
		if count > 0 && user.ID != eid.ID {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   strconv.FormatBool(true),
				"data":    "",
				"message": "The email fild must be unique",
			})
		} else {
			if err = res.Update(user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errors.CheckErr(err)
			} else {
				data, _ := json.Marshal(user)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   strconv.FormatBool(false),
					"data":    string(data),
					"message": "User updated",
				})
			}
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var user models.User
	res := models.UsersModel.Find(id)
	err := res.One(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   strconv.FormatBool(false),
			"data":    "",
			"message": "user not found",
		})
	} else {
		_ = json.NewDecoder(r.Body).Decode(&user)

		if err = res.Delete(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errors.CheckErr(err)
		} else {
			data, _ := json.Marshal(user)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   strconv.FormatBool(false),
				"data":    string(data),
				"message": "User deleted",
			})
		}
	}
}
