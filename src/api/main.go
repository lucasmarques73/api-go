package main

import (
	"api/Controllers"
	"fmt"
	"html/template"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	routes := mux.NewRouter().StrictSlash(true)
	routesProtected := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", home)
	routes.HandleFunc("/auth/login", Controllers.Login).Methods("POST")

	routesProtected.HandleFunc("/users", Controllers.GetAllUsers).Methods("GET")
	routesProtected.HandleFunc("/users", Controllers.CreateUser).Methods("POST")
	routesProtected.HandleFunc("/users/{id}", Controllers.GetUser).Methods("GET")
	routesProtected.HandleFunc("/users/{id}", Controllers.UpdateUser).Methods("PUT")
	routesProtected.HandleFunc("/users/{id}", Controllers.DeleteUser).Methods("DELETE")

	routesProtected.HandleFunc("/widgets", Controllers.GetAllWidgets).Methods("GET")
	routesProtected.HandleFunc("/widgets", Controllers.CreateWidget).Methods("POST")
	routesProtected.HandleFunc("/widgets/{id}", Controllers.GetWidget).Methods("GET")
	routesProtected.HandleFunc("/widgets/{id}", Controllers.UpdateWidget).Methods("PUT")
	routesProtected.HandleFunc("/widgets/{id}", Controllers.DeleteWidget).Methods("DELETE")

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodRS512,
	})

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(routesProtected))

	routes.PathPrefix("/").Handler(an)

	n := negroni.Classic()
	n.UseHandler(routes)
	n.Use(negroni.NewLogger())

	port := ":80"
	n.Run(port)

	fmt.Println("Server running in port:", port)
	// log.Fatal(http.ListenAndServe(port, routes))
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("Views/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
