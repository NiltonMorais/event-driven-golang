package queue

import "context"

type Publisher interface {
	Publish(ctx context.Context, eventPayload interface{})
}
