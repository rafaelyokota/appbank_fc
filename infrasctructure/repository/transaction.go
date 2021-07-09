package repository

import (
	"database/sql"
	"errors"
	"github.com/rafaelyokota/codebank/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func (t *TransactionRepositoryDb) GetCc(cc domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard
	stmt, err := t.db.Prepare("select id, balance, balance_limit from credit_cards where number=$1")
	if err != nil {
		return c, err
	}
	if err = stmt.QueryRow(cc.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("credit card does not exists")
	}
	return c, nil
}

func (t *TransactionRepositoryDb) CreateCc(creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(`insert into credit_cards(id, name, number, expiration_month,expiration_year, CVV,balance, balance_limit) 
								values($1,$2,$3,$4,$5,$6,$7,$8)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpireMonth,
		creditCard.ExpireYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb{
	return &TransactionRepositoryDb{db:db}
}

func (t *TransactionRepositoryDb)SaveTransaction(transacao domain.Transaction, cc domain.CreditCard) error{

	stmt, err := t.db.Prepare( `insert into trasaction(id, credit_card_id, amount, status, description, store, create_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`)
		if nil != err{
			return err
		}
		_, err = stmt.Exec(
			transacao.ID,
			transacao.CreditCardId,
			transacao.Amount,
			transacao.Status,
			transacao.Description,
			transacao.Store,
			transacao.CreateAt,
			)

		if nil != err{
			return err
		}

		if transacao.Status == "approved"{
			err := t.updateBalanace(cc)
			if err != nil{
				return err
			}
		}
		err = stmt.Close()

		if nil != err{
			return err
		}

		return nil
}

func (t *TransactionRepositoryDb) updateBalanace(cc domain.CreditCard) error{
	_, err := t.db.Exec("update credit_cards set balance = $1 where id = $2",
		cc.Balance,
		cc.ID)
	if err != nil{
		return err
	}
	return nil
}

func (t *TransactionRepositoryDb) GetCreditCard(cc domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard
	stmt, err := t.db.Prepare("select id, balance, balance_limit from credit_cards where number=$1")
	if err != nil {
		return c, err
	}
	if err = stmt.QueryRow(cc.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("credit card does not exists")
	}
	return c, nil
}
