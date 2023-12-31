package main

import (

	// "html/template"
	"net/http"

	// "github.com/gorilla/mux"
	"admin-panel/controller"
	"admin-panel/middleware"
	"admin-panel/model"

	"github.com/gorilla/mux"
	// "github.com/lib/pq"
	// "example.com/jwt-demo/controller"
)

func main() {
	model.Connect()
	r := mux.NewRouter()

	indexHandler := http.HandlerFunc(controller.LoginHandler)
	adminHandler := http.HandlerFunc(controller.AdminPanel)
	signupHandler := http.HandlerFunc(controller.SignupHandler)

	r.Handle("/signup", middleware.Auth(signupHandler)).Methods("GET")
	// r.HandleFunc("/signup",controller.SignupHandler).Methods("GET")

	r.HandleFunc("/signup", controller.SignUp).Methods("POST")

	r.Handle("/", middleware.Auth(indexHandler)).Methods("GET")

	r.HandleFunc("/", controller.Login).Methods("POST")

	r.HandleFunc("/home", controller.HomeHandler).Methods("GET")

	r.Handle("/adminpanel", middleware.AdminAuth(adminHandler)).Methods("GET")
	// r.HandleFunc("/adminpanel",controller.AdminPanel).Methods("GET")
	r.HandleFunc("/adminpanel", controller.AddUser).Methods("POST")

	r.HandleFunc("/delete", controller.DeleteUser).Methods("GET")

	r.HandleFunc("/update", controller.UpdateUser).Methods("GET")
	r.HandleFunc("/update", controller.UpdateUserReal).Methods("POST")

	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	http.ListenAndServe(":8080", r)

}
