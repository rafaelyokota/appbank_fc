package usecase

import (
	"github.com/rafaelyokota/codebank/domain"
	"github.com/rafaelyokota/codebank/dto"
	"time"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUserCaseTransaction(transacao domain.TransactionRepository)*UseCaseTransaction{
	return &UseCaseTransaction{TransactionRepository: transacao}
}

func (u *UseCaseTransaction)ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error){
	cc := u.hydrateCc(transactionDto)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCc(*cc)
	if err!= nil{
		return domain.Transaction{}, err
	}
	cc.ID = ccBalanceAndLimit.ID
	cc.Name = ccBalanceAndLimit.Name
	cc.Limit = ccBalanceAndLimit.Limit
	cc.Balance = ccBalanceAndLimit.Balance

	t := u.newTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAndValidate(cc)

	err = u.TransactionRepository.SaveTransaction(*t, *cc)

	if nil != err{
		return domain.Transaction{}, err
	}

	return *t, nil
}

func (*UseCaseTransaction) hydrateCc(transactionDto dto.Transaction) *domain.CreditCard{
	cc := domain.NewCreditCard()
	cc.Name = transactionDto.Name
	cc.Number = transactionDto.Number
	cc.ExpireMonth = transactionDto.ExpireMonth
	cc.ExpireYear = transactionDto.ExpireYear
	cc.CVV = transactionDto.CVV
	return cc
}

func (*UseCaseTransaction) newTransaction(transaction dto.Transaction, cc domain.CreditCard) *domain.Transaction{
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transaction.Amount
	t.Store = transaction.Store
	t.Description = transaction.Description
	t.CreateAt = time.Now()
	return t
}