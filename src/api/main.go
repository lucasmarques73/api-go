package main

import (
	"api/Controllers"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", home)
	routes.HandleFunc("/users", Controllers.GetAllUsers).Methods("GET")
	routes.HandleFunc("/users", Controllers.CreateUser).Methods("POST")
	routes.HandleFunc("/users/{id}", Controllers.GetUser).Methods("GET")
	routes.HandleFunc("/users/{id}", Controllers.UpdateUser).Methods("PUT")
	routes.HandleFunc("/users/{id}", Controllers.DeleteUser).Methods("DELETE")

	routes.HandleFunc("/widgets", Controllers.GetAllWidgets).Methods("GET")
	routes.HandleFunc("/widgets", Controllers.CreateWidget).Methods("POST")
	routes.HandleFunc("/widgets/{id}", Controllers.GetWidget).Methods("GET")
	routes.HandleFunc("/widgets/{id}", Controllers.UpdateWidget).Methods("PUT")
	routes.HandleFunc("/widgets/{id}", Controllers.DeleteWidget).Methods("DELETE")

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
