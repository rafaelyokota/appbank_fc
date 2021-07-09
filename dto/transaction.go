package dto

import "time"

type Transaction struct {
	ID string
	Name string
	Number string
	ExpireMonth int32
	ExpireYear int32
	CVV int32
	Amount float64
	Store string
	Description string
	CreatAt time.Time
}