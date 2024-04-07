package factory

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/NiltonMorais/event-driven-golang/internal/application/controller"
	"github.com/NiltonMorais/event-driven-golang/internal/application/handler"
	"github.com/NiltonMorais/event-driven-golang/internal/application/usecase"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/queue"
	infraQueue "github.com/NiltonMorais/event-driven-golang/internal/infra/queue"
)

type Application struct {
	queue                      queue.Queue
	userController             *controller.UserController
	orderController            *controller.OrderController
	sendWelcomeEmailHandler    *handler.SendWelcomeEmailHandler
	processOrderPaymentHandler *handler.ProcessOrderPaymentHandler
	stockMovementHandller      *handler.StockMovementHandler
	sendOrderEmailHandler      *handler.SendOrderEmailHandler
}

func NewApplication() (*Application, error) {
	//queue := infraQueue.NewMemoryQueueAdapter()
	queueUri := os.Getenv("QUEUE_URI")
	queue := infraQueue.NewRabbitMQAdapter(queueUri)

	sendWelcomeEmailHandler := handler.NewSendWelcomeEmailHandler(queue)
	processOrderPaymentHandler := handler.NewProcessOrderPaymentHandler()
	stockMovementHandller := handler.NewStockMovementHandler()
	sendOrderEmailHandler := handler.NewSendOrderEmailHandler()

	createUserUseCase := usecase.NewCreateUserUseCase(queue)
	userController := controller.NewUserController(createUserUseCase)

	createOrderUseCase := usecase.NewCreateOrderUseCase(queue)
	orderController := controller.NewOrderController(createOrderUseCase)

	return &Application{
		queue:                      queue,
		userController:             userController,
		orderController:            orderController,
		sendWelcomeEmailHandler:    sendWelcomeEmailHandler,
		processOrderPaymentHandler: processOrderPaymentHandler,
		stockMovementHandller:      stockMovementHandller,
		sendOrderEmailHandler:      sendOrderEmailHandler,
	}, nil
}

func (app *Application) Run(ctx context.Context) error {
	err := app.queue.Connect(ctx)
	if err != nil {
		return err
	}
	defer app.queue.Disconnect(ctx)
	log.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
