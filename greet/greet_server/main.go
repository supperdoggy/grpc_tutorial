package main

import (
	"github.com/supperdoggy/grpc_course/greet/greet_server/internal/service"
	"github.com/supperdoggy/grpc_course/greet/greetpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger, _ := zap.NewDevelopment()

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatal("error listening to adress", zap.Error(err))
	}

	server := service.NewServer(logger)

	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcServer, server)

	logger.Info("serving the server", zap.Any("listener", listener.Addr()))

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("error serving server", zap.Error(err))
	}

}
