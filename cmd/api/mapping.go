package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/infrastructure"
	"github.com/gorilla/mux"
)

func ConfigureMappings(app mux.Router, handlers infrastructure.HandlerContainer) {

	//material
	app.HandleFunc("/api/material", handlers.GetAllMaterial.Handle).Methods("GET")
	app.HandleFunc("/api/material", handlers.CreateMaterial.Handle).Methods("POST")

	fmt.Println("Starting app in localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", &app))

}
