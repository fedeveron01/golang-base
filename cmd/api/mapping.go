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
	app.HandleFunc("/api/material", handlers.MaterialHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/material", handlers.MaterialHandler.Create).Methods("POST")

	//material type
	app.HandleFunc("/api/materialType", handlers.MaterialTypeHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/materialType", handlers.MaterialTypeHandler.Create).Methods("POST")
	app.HandleFunc("/api/materialType/units", handlers.MaterialTypeHandler.GetUnitsOfMeasurement).Methods("GET")

	//user
	app.HandleFunc("/api/user/signup", handlers.UserHandler.Signup).Methods("POST")
	app.HandleFunc("/api/user/login", handlers.UserHandler.Login).Methods("POST")
	app.HandleFunc("/api/user/logout", handlers.UserHandler.Logout).Methods("POST")

	//employee
	app.HandleFunc("/api/employee", handlers.EmployeeHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/employee/{id}", handlers.EmployeeHandler.GetById).Methods("GET")

	//charge
	app.HandleFunc("/api/charge", handlers.ChargeHandler.GetAll).Methods("GET")
	app.HandleFunc("/api/charge", handlers.ChargeHandler.Create).Methods("POST")

	fmt.Println("Starting app in localhost:8080")
	handler := cors.AllowAll().Handler(&app)
	log.Fatal(http.ListenAndServe(":8080", handler))

}
