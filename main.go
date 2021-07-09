package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	domain "github.com/rafaelyokota/codebank/domain"
	"github.com/rafaelyokota/codebank/infrasctructure/repository"
	"github.com/rafaelyokota/codebank/usecase"
	"log"
)

func main(){
	fmt.Println("Hello")
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "rafael"
	cc.ExpireYear = 2021
	cc.ExpireMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCc(*cc)
	if err != nil{
		fmt.Println(err)
	}
}
func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction{
		transactionRepo := repository.NewTransactionRepositoryDb(db)
		usecase := usecase.NewUserCaseTransaction(transactionRepo)
		return *usecase

}

func setupDb()*sql.DB{
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil{
		log.Fatal("error connection to database")
	}
	return db
}