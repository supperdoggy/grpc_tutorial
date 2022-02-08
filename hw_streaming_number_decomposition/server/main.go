package main

import (
	"github.com/supperdoggy/grpc_course/hw_streaming_number_decomposition/NumberDecompositionpb"
	service2 "github.com/supperdoggy/grpc_course/hw_streaming_number_decomposition/server/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger, _ := zap.NewDevelopment()
	service := service2.NewService(logger)

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatal("error listening", zap.Error(err))
	}

	server := grpc.NewServer()
	NumberDecompositionpb.RegisterStreamGreetingServiceServer(server, service)

	if err := server.Serve(listener); err != nil {
		logger.Fatal("error serving service", zap.Error(err))
	}

}