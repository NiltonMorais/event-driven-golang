package factory

import (
	"net/http"
	"reflect"

	orderEvent "github.com/NiltonMorais/event-driven-golang/internal/order/domain/event"
	userEvent "github.com/NiltonMorais/event-driven-golang/internal/user/domain/event"
)

func RegisterConsumers(app *Application) {
	var list map[reflect.Type][]func(w http.ResponseWriter, r *http.Request) = map[reflect.Type][]func(w http.ResponseWriter, r *http.Request){
		reflect.TypeOf(userEvent.UserRegisteredEvent{}): {
			app.userController.SendWelcomeEmail,
		},
		reflect.TypeOf(orderEvent.OrderCreatedEvent{}): {
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
