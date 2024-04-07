package usecase

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/user/domain/event"
	"github.com/NiltonMorais/event-driven-golang/internal/user/domain/queue"
)

type SendWelcomeEmailUseCase struct {
	publisher queue.Publisher
}

func NewSendWelcomeEmailUseCase(publisher queue.Publisher) *SendWelcomeEmailUseCase {
	return &SendWelcomeEmailUseCase{
		publisher: publisher,
	}
}

func (h *SendWelcomeEmailUseCase) Execute(ctx context.Context, input *event.UserRegisteredEvent) error {
	fmt.Println("--- SendWelcomeEmailUseCase ---")
	fmt.Printf("--- MAIL to %s: Welcome %s --- \n", input.Email, input.Name)
	h.publisher.Publish(ctx, event.WelcomeEmailSentEvent{Email: input.Email})
	return nil
}
