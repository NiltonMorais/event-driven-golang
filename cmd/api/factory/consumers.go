package factory

import (
	"net/http"
	"reflect"

	"github.com/NiltonMorais/event-driven-golang/internal/order/domain/event"
)

func RegisterConsumers(app *Application) {
	var list map[reflect.Type][]func(w http.ResponseWriter, r *http.Request) = map[reflect.Type][]func(w http.ResponseWriter, r *http.Request){
		reflect.TypeOf(event.UserRegisteredEvent{}): {
			app.userController.SendWelcomeEmail,
		},
		reflect.TypeOf(event.OrderCreatedEvent{}): {
			app.orderController.ProcessOrderPayment,
			app.orderController.StockMovement,
			app.orderController.SendOrderEmail,
		},
	}

	for eventType, handlers := range list {
		for _, handler := range handlers {
			app.queue.ListenerRegister(eventType, handler)
		}
	}
}
