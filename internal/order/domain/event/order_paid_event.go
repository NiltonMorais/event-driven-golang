package event

import (
	"time"
)

type OrderPaidEvent struct {
	OrderId     string
	PaidValue   float64
	PaymentDate time.Time
}
