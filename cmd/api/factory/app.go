package factory

import (
	"context"
	"log"
	"net/http"
	"os"
	"reflect"

	orderController "github.com/NiltonMorais/event-driven-golang/internal/order/application/controller"
	orderUsecase "github.com/NiltonMorais/event-driven-golang/internal/order/application/usecase"
	orderEvent "github.com/NiltonMorais/event-driven-golang/internal/order/domain/event"
	userController "github.com/NiltonMorais/event-driven-golang/internal/user/application/controller"
	userUsecase "github.com/NiltonMorais/event-driven-golang/internal/user/application/usecase"
	userEvent "github.com/NiltonMorais/event-driven-golang/internal/user/domain/event"
	"github.com/NiltonMorais/event-driven-golang/pkg/queue"
)

type Application struct {
	queue           queue.Queue
	userController  *userController.UserController
	orderController *orderController.OrderController
}

func NewApplication() (*Application, error) {
	// queue := queue.NewMemoryQueueAdapter()
	queueUri := os.Getenv("QUEUE_URI")
	queue := queue.NewRabbitMQAdapter(queueUri)

	createUserUseCase := userUsecase.NewCreateUserUseCase(queue)
	sendWelcomeEmailUseCase := userUsecase.NewSendWelcomeEmailUseCase(queue)
	userController := userController.NewUserController(createUserUseCase, sendWelcomeEmailUseCase)

	createOrderUseCase := orderUsecase.NewCreateOrderUseCase(queue)
	processOrderPaymentUseCase := orderUsecase.NewProcessOrderPaymentUseCase()
	stockMovementUseCase := orderUsecase.NewStockMovementUseCase()
	sendOrderEmailUseCase := orderUsecase.NewSendOrderEmailUseCase()
	orderController := orderController.NewOrderController(createOrderUseCase, processOrderPaymentUseCase, stockMovementUseCase, sendOrderEmailUseCase)

	return &Application{
		queue:           queue,
		userController:  userController,
		orderController: orderController,
	}, nil
}

func (app *Application) RunServer(ctx context.Context) error {
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

func (app *Application) StartConsumingQueues(ctx context.Context) error {
	err := app.queue.Connect(ctx)
	if err != nil {
		return err
	}

	OrderCreatedEvent := reflect.TypeOf(orderEvent.OrderCreatedEvent{}).Name()
	UserRegisteredEvent := reflect.TypeOf(userEvent.UserRegisteredEvent{}).Name()

	go func(ctx context.Context, queueName string) {
		err = app.queue.StartConsuming(ctx, queueName)
		if err != nil {
			log.Fatalf("Error running consumer %s: %s", queueName, err)
		}
	}(ctx, OrderCreatedEvent)

	go func(ctx context.Context, queueName string) {
		err = app.queue.StartConsuming(ctx, queueName)
		if err != nil {
			log.Fatalf("Error running consumer %s: %s", queueName, err)
		}
	}(ctx, UserRegisteredEvent)

	return nil
}

func (app *Application) DisconnectQueue(ctx context.Context) error {
	err := app.queue.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}
