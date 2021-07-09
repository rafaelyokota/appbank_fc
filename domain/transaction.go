package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type TransactionRepository interface{
	SaveTransaction(transacao Transaction, cc CreditCard) error
	GetCc(cc CreditCard)(CreditCard, error)
	CreateCc(cc CreditCard) error
}

type Transaction struct {
	ID string
	Amount float64
	Status string
	Description string
	Store string
	CreditCardId string
	CreateAt time.Time

}

func NewTransaction() *Transaction{
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreateAt = time.Now()

	return t
}

func (t *Transaction)ProcessAndValidate(cc *CreditCard){
	if t.Amount + cc.Balance > cc.Limit{
		t.Status = "rejected"
	}else{
		t.Status = "approved"
		cc.Balance = cc.Balance + t.Amount
	}
}