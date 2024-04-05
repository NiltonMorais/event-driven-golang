package factory

import (
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/internal/application/handler"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
)

func RegisterHandlers(app *Application) {
	var list map[reflect.Type][]handler.Handler = map[reflect.Type][]handler.Handler{
		reflect.TypeOf(event.UserRegisteredEvent{}): {
			app.sendWelcomeEmailHandler,
		},
		reflect.TypeOf(event.OrderCreatedEvent{}): {
			app.processOrderPaymentHandler,
			app.stockMovementHandller,
			app.sendOrderEmailHandler,
		},
	}

	for eventType, handlers := range list {
		for _, handler := range handlers {
			app.queue.Register(eventType, handler.Execute)
		}
	}
}
