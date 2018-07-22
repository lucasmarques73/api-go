package controllers

import (
	"api/errors"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	u, err := models.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found")
	}
	json.NewEncoder(w).Encode(u)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]
	u, err := models.GetUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found")
	} else {
		json.NewEncoder(w).Encode(u)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	if _, err := models.UsersModel.Insert(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errors.CheckErr(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found")
	} else {
		_ = json.NewDecoder(r.Body).Decode(&u)

		u, err = models.UpdateUser(u)
		errors.CheckErr(err)

		json.NewEncoder(w).Encode(u)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]

	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found")
	} else {
		err := models.DeleteUser(u)
		if err != nil {
			json.NewEncoder(w).Encode("Deleted")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		}

	}
}
