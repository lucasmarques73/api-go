package main

import (
	"api/controllers"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", home)
	routes.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	routes.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	routes.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	routes.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	routes.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// routes.HandleFunc("/widgets", controllers.getAllWidgets).Methods("GET")
	// routes.HandleFunc("/widgets", controllers.createWidget).Methods("POST")
	// routes.HandleFunc("/widgets/{id}", controllers.getWidget).Methods("GET")
	// routes.HandleFunc("/widgets/{id}", controllers.getUser).Methods("PUT")
	// routes.HandleFunc("/widgets/{id}", controllers.getUser).Methods("DELETE")

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
