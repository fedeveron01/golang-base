package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fedeveron01/golang-base/cmd/infrastructure"
	"github.com/gorilla/mux"
)

func ConfigureMappings(app mux.Router, handlers infrastructure.HandlerContainer) {
	app.HandleFunc("/api", handlers.CalculateAge.Handle).Methods("GET")
	app.HandleFunc("/server", handlers.CalculateAge.Handle).Methods("GET")
	fmt.Println("Starting app in localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", &app))

}
