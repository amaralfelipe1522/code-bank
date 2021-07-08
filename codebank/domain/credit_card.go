package	domain

import "time"

type CreditCard struct {
	ID string
	Name string
	Number string
	ExpirationMonth int32
	ExpirationYear int32
	CVV int32
	Balance float64
	Limit float64
	CreatedAt time.Time
} // 55min