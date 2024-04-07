package queue

import "context"

type Publisher interface {
	Publish(ctx context.Context, body interface{}) error
}
