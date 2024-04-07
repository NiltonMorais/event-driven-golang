package main

import (
	"context"
	"log"
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/cmd/consumer/factory"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

func main() {
	app, err := factory.NewApplication()
	if err != nil {
		log.Fatalf("Error creating application: %s", err)
	}
	factory.RegisterHandlers(app)

	ctx := context.Background()
	OrderCreatedEvent := reflect.TypeOf(event.OrderCreatedEvent{}).Name()
	UserRegisteredEvent := reflect.TypeOf(event.UserRegisteredEvent{}).Name()

	go func(ctx context.Context, queueName string) {
		err = app.StartConsuming(ctx, queueName)
		if err != nil {
			log.Fatalf("Error running consumer %s: %s", queueName, err)
		}
	}(ctx, OrderCreatedEvent)

	go func(ctx context.Context, queueName string) {
		err = app.StartConsuming(ctx, queueName)
		if err != nil {
			log.Fatalf("Error running consumer %s: %s", queueName, err)
		}
	}(ctx, UserRegisteredEvent)

	var forever chan struct{}
	<-forever
}
