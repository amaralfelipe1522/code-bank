package server

import (
	"github.com/amaralfelipe1522/codebank/usecase"
	"github.com/amaralfelipe1522/codebank/infrastructure/grpc/service"
	"github.com/amaralfelipe1522/codebank/infrastructure/grpc/pb"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"log"
	"net"
)
type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer{
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("could not listen tcp port")
	}

	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)

}

