package factory

import "net/http"

func ResgisterRoutes(app *Application) {
	http.HandleFunc("/", app.userController.HelloWorld)
	http.HandleFunc("POST /create-user", app.userController.CreateUser)
	http.HandleFunc("POST /create-order", app.orderController.CreateOrder)
}
