package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreditCard struct {
	ID string
	Name string
	Number string
	ExpireMonth int32
	ExpireYear int32
	CVV int32
	Balance float64
	Limit float64
	CreatAt time.Time

}

func NewCreditCard() *CreditCard{
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatAt = time.Now()
	return c
}