package queue

import (
	"context"
	"net/http"
	"reflect"
)

type Queue interface {
	ListenerRegister(eventType reflect.Type, handler func(w http.ResponseWriter, r *http.Request))
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Publish(ctx context.Context, body interface{}) error
	StartConsuming(ctx context.Context, queueName string) error
}
