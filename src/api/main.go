package main

import (
	"api/users"
	"api/widgets"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", home)
	routes.HandleFunc("/users", getAllUsers).Methods("GET")
	// routes.HandleFunc("/users", createUser).Methods("POST")
	routes.HandleFunc("/users/{id}", getUser).Methods("GET")
	// routes.HandleFunc("/users/{id}", getUser).Methods("PUT")
	// routes.HandleFunc("/users/{id}", getUser).Methods("DELETE")

	routes.HandleFunc("/widgets", getAllWidgets).Methods("GET")
	// routes.HandleFunc("/widgets", createWidget).Methods("POST")
	routes.HandleFunc("/widgets/{id}", getWidget).Methods("GET")
	// routes.HandleFunc("/widgets/{id}", getUser).Methods("PUT")
	// routes.HandleFunc("/widgets/{id}", getUser).Methods("DELETE")

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

// UsersController
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(users.ListUsers())
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(users.GetUserById(id))
}

// WidgetsController
func getAllWidgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(widgets.ListWidgets())
}

func getWidget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(widgets.GetWidgetById(id))
}
