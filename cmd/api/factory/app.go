package factory

import (
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
	queue := infraQueue.NewMemoryQueue()

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
