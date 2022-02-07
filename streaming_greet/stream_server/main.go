package main

import (
	"github.com/supperdoggy/grpc_course/streaming_greet/GreetStreampb"
	service2 "github.com/supperdoggy/grpc_course/streaming_greet/stream_server/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger, _ := zap.NewDevelopment()
	service := service2.NewServer(logger)

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatal("error listening", zap.Error(err))
	}

	server := grpc.NewServer()
	GreetStreampb.RegisterStreamGreetingServiceServer(server, service)

	if err := server.Serve(listener); err != nil {
		logger.Fatal("error serving server", zap.Error(err))
	}
}
