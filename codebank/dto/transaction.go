package dto

import "time"

// Responsável por receber e preparar os dados externos
type Transaction struct {
	ID string
	Name string
	Number string
	ExpirationMonth int32
	ExpirationMonth int32
	CVV int32
	Amount float64
	Store string
	Description string
	CreatedAt time.Time
}