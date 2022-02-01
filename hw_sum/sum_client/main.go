package main

import (
	"context"
	"github.com/supperdoggy/grpc_course/hw_sum/sumpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewDevelopment()

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("error dialing with server", zap.Error(err))
	}
	defer func () {
		err := conn.Close()
		if err != nil {
			logger.Error("error closing connection", zap.Error(err))
		}
	}()

	service := sumpb.NewSumServiceClient(conn)

	resp, err := service.Add(context.Background(), &sumpb.AddRequest{A: -1, B: 3})
	if err != nil {
		logger.Fatal("got nil resp", zap.Error(err))
	}
	logger.Info("got resp", zap.Any("resp", resp.GetSum()))

}
