package event

import (
	"encoding/json"
	"reflect"
	"time"
)

type DomainEvent struct {
	Type    reflect.Type
	Date    time.Time
	Payload []byte
}

func GetPayload[T any](e DomainEvent) (*T, error) {
	var payload T
	err := json.Unmarshal(e.Payload, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
