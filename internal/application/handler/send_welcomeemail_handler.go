package handler

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/queue"
)

type SendWelcomeEmailHandler struct {
	publisher queue.Publisher
}

func NewSendWelcomeEmailHandler(publisher queue.Publisher) *SendWelcomeEmailHandler {
	return &SendWelcomeEmailHandler{
		publisher: publisher,
	}
}

func (h *SendWelcomeEmailHandler) Execute(ctx context.Context, e event.DomainEvent) error {
	payload, err := event.GetPayload[event.UserRegisteredEvent](e)
	if err != nil {
		return err
	}
	fmt.Println("--- SendWelcomeEmailHandler ---")
	fmt.Printf("--- MAIL to %s: Welcome %s --- \n", payload.Email, payload.Name)

	h.publisher.Publish(ctx, event.WelcomeEmailSentEvent{Email: payload.Email})
	return nil
}
