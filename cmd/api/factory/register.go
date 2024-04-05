package factory

import (
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

func RegisterHandlers(app *Application) {
	app.queue.Register(reflect.TypeOf(event.UserRegisteredEvent{}), app.sendWelcomeEmailHandler.Execute)
	app.queue.Register(reflect.TypeOf(event.OrderCreatedEvent{}), app.processOrderPaymentHandler.Execute)
	app.queue.Register(reflect.TypeOf(event.OrderCreatedEvent{}), app.stockMovementHandller.Execute)
	app.queue.Register(reflect.TypeOf(event.OrderCreatedEvent{}), app.sendOrderEmailHandler.Execute)
}
