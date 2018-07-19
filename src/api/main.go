package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"html/template"
	"log"
	"net/http"	
)

func main(){
	routes := mux.NewRouter()

	routes.HandleFunc("/", HomeHandler)

	port := ":80"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}


func HomeHandler (w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("views/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}