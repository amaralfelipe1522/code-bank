package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/amaralfelipe1522/codebank/domain"
	"github.com/amaralfelipe1522/codebank/usecase"
	"github.com/amaralfelipe1522/codebank/infrastructure/repository"
)

func main() {
	db:= setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()

	cc.Number = "123456"
	cc.Name = "Felipe Amaral"
	cc.ExpirationMonth = 8
	cc.ExpirationYear = 2028
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)

	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	"db",
	"5432",
	"postgres",
	"root",
	"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}