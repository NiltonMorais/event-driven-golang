package main

import (
	"log"
	"net/http"

	"github.com/NiltonMorais/event-driven-golang/cmd/api/factory"
)

func main() {
	app, err := factory.NewApplication()
	if err != nil {
		log.Fatalf("Error creating application: %s", err)
	}
	factory.RegisterHandlers(app)
	factory.RouteApplication(app)
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
