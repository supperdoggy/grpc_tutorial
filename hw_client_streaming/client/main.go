package main

import (
	"context"
	"github.com/supperdoggy/grpc_course/hw_client_streaming/avrgpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewDevelopment()

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("error dialing", zap.Error(err))
	}

	defer func () {
		err := conn.Close()
		if err != nil {
			logger.Error("error closing connection", zap.Error(err))
		}
	}()

	service := avrgpb.NewStreamGreetingServiceClient(conn)

	stream, err := service.AvrgClientNumberStream(context.Background(), )
	if err != nil {
		logger.Fatal("error stream", zap.Error(err))
	}

	for _, v := range []int64{1, 2, 3, 4, 5} {
		err := stream.Send(&avrgpb.AvrgClientNumberStreamRequest{Number: v})
		if err != nil {
			logger.Fatal("error", zap.Error(err))
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		logger.Fatal("error", zap.Error(err))
	}

	logger.Info("resp", zap.Any("resp", resp))

}
