package usecase

import (
	"context"
	"fmt"

	"github.com/NiltonMorais/event-driven-golang/internal/domain/entity"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/event"
	"github.com/NiltonMorais/event-driven-golang/internal/domain/queue"
)

type CreateUserUseCase struct {
	publisher queue.Publisher
}

func NewCreateUserUseCase(publisher queue.Publisher) *CreateUserUseCase {
	return &CreateUserUseCase{
		publisher: publisher,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, name, email string) error {
	fmt.Println("--- CreateUserUseCase ---")
	user, err := entity.NewUserEntity(name, email)
	if err != nil {
		return err
	}
	event := event.UserRegisteredEvent{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
	u.publisher.Publish(ctx, event)
	return nil
}
