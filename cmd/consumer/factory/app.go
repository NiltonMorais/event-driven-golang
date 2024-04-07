package factory

import (
	"context"
	"os"

	"github.com/NiltonMorais/event-driven-golang/internal/application/handler"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/queue"
	infraQueue "github.com/NiltonMorais/event-driven-golang/internal/infra/queue"
)

type Application struct {
	queue                      queue.Queue
	sendWelcomeEmailHandler    *handler.SendWelcomeEmailHandler
	processOrderPaymentHandler *handler.ProcessOrderPaymentHandler
	stockMovementHandller      *handler.StockMovementHandler
	sendOrderEmailHandler      *handler.SendOrderEmailHandler
}

func NewApplication() (*Application, error) {
	queueUri := os.Getenv("QUEUE_URI")
	queue := infraQueue.NewRabbitMQAdapter(queueUri)

	sendWelcomeEmailHandler := handler.NewSendWelcomeEmailHandler(queue)
	processOrderPaymentHandler := handler.NewProcessOrderPaymentHandler()
	stockMovementHandller := handler.NewStockMovementHandler()
	sendOrderEmailHandler := handler.NewSendOrderEmailHandler()

	return &Application{
		queue:                      queue,
		sendWelcomeEmailHandler:    sendWelcomeEmailHandler,
		processOrderPaymentHandler: processOrderPaymentHandler,
		stockMovementHandller:      stockMovementHandller,
		sendOrderEmailHandler:      sendOrderEmailHandler,
	}, nil
}

func (app *Application) StartConsuming(ctx context.Context, queueName string) error {
	err := app.queue.Connect(ctx)
	if err != nil {
		return err
	}
	err = app.queue.StartConsuming(ctx, queueName)
	if err != nil {
		return err
	}
	defer app.queue.Disconnect(ctx)
	return nil
}
