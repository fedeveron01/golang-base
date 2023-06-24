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
	app.HandleFunc("/whatsappServer", handlers.Whatsapp.Handle).Methods("GET")
	app.HandleFunc("/api/subscriptions", handlers.GetAllSubscription.Handle).Methods("GET")
	app.HandleFunc("/api/subscriptions/{id}", handlers.GetSubscription.Handle).Methods("GET")
	app.HandleFunc("/api/subscriptions/{id}", handlers.EditSubscription.Handle).Methods("PUT")
	app.HandleFunc("/api/subscriptions/{id}", handlers.DeleteSubscription.Handle).Methods("DELETE")
	app.HandleFunc("/api/subscriptions", handlers.CreateSubscription.Handle).Methods("POST")

	fmt.Println("Starting app in localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", &app))

}
