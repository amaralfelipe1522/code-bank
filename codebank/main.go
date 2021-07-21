package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/amaralfelipe1522/codebank/usecase"
	"github.com/amaralfelipe1522/codebank/infrastructure/repository"
	"github.com/amaralfelipe1522/codebank/infrastructure/kafka"
	"github.com/amaralfelipe1522/codebank/infrastructure/grpc/server"
)

func main() {
	db:= setupDb()
	defer db.Close()

	// Código de teste para criação do cartão de crédito
	// cc := domain.NewCreditCard()
	// cc.Number = "123456"
	// cc.Name = "Felipe Amaral"
	// cc.ExpirationMonth = 8
	// cc.ExpirationYear = 2028
	// cc.CVV = 123
	// cc.Limit = 1000
	// cc.Balance = 0
	// repo := repository.NewTransactionRepositoryDb(db)
	// err := repo.CreateCreditCard(*cc)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer("host.docker.internal:9094")
	return producer
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

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Rodando gRPC Server")
	grpcServer.Serve()
}