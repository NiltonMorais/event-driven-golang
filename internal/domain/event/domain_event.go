package event

import (
	"reflect"
	"time"
)

type DomainEvent struct {
	Type    reflect.Type
	Date    time.Time
	Payload interface{}
}

func GetPayload[T any](e DomainEvent) T {
	return e.Payload.(T)
}
