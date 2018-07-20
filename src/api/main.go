package main

import (
	"api/errors"
	"api/users"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", home)
	routes.HandleFunc("/users", getAllUsers).Methods("GET")
	routes.HandleFunc("/users", createUser).Methods("POST")
	routes.HandleFunc("/users/{id}", getUser).Methods("GET")
	// routes.HandleFunc("/users/{id}", getUser).Methods("PUT")
	// routes.HandleFunc("/users/{id}", getUser).Methods("DELETE")

	port := ":80"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(users.ListUsers())
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var u users.User

	body, err := ioutil.ReadAll(r.Body)
	errors.CheckErr(err)

	err = r.Body.Close()
	errors.CheckErr(err)

	json.Unmarshal(body, &u)

	u = users.CreateUser(u)

	json.NewEncoder(w).Encode(u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(users.GetUserById(id))
}
