package event

import "github.com/NiltonMorais/event-driven-golang/internal/domain/entity"

type OrderCreatedEvent struct {
	Order *entity.OrderEntity
}
