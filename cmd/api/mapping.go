package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/infrastructure"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func ConfigureMappings(app mux.Router, handlers infrastructure.HandlerContainer) {
	//ping
	app.HandleFunc("/ping", handlers.Ping.Handle).Methods("GET")

	//material
	app.HandleFunc("/api/material", handlers.GetAllMaterial.Handle).Methods("GET")
	app.HandleFunc("/api/material", handlers.CreateMaterial.Handle).Methods("POST")

	//user
	app.HandleFunc("/api/user/signup", handlers.CreateUser.Handle).Methods("POST")
	app.HandleFunc("/api/user/login", handlers.LoginUser.Handle).Methods("POST")
	app.HandleFunc("/api/user/logout", handlers.LogoutUser.Handle).Methods("POST")

	//employee
	app.HandleFunc("/api/employee", handlers.GetAllEmployee.Handle).Methods("GET")

	fmt.Println("Starting app in localhost:8080")
	handler := cors.AllowAll().Handler(&app)
	log.Fatal(http.ListenAndServe(":8080", handler))

}
