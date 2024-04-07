package queue

import (
	"net/http"
	"reflect"
)

type Listener struct {
	eventType reflect.Type
	callback  func(w http.ResponseWriter, r *http.Request)
}
