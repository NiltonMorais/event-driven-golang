package main

import (
	"context"
	"log"

	"github.com/NiltonMorais/event-driven-golang/cmd/api/factory"
)

func main() {
	app, err := factory.NewApplication()
	if err != nil {
		log.Fatalf("Error creating application: %s", err)
	}
	factory.RouteApplication(app)

	ctx := context.Background()
	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("Error running application: %s", err)
	}
}
