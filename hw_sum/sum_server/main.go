package main

import (
	service2 "github.com/supperdoggy/grpc_course/hw_sum/sum_server/internal/service"
	"github.com/supperdoggy/grpc_course/hw_sum/sumpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger, _ := zap.NewDevelopment()
	service := service2.NewService(logger)

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatal("error starting listener", zap.Error(err))
	}

	server := grpc.NewServer()
	sumpb.RegisterSumServiceServer(server, service)

	logger.Info("starting server", zap.Any("addr", listener.Addr()))
	if err := server.Serve(listener); err != nil {
		logger.Fatal("serving error", zap.Error(err))
	}
}
