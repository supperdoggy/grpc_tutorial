package main

import (
	"context"
	"github.com/supperdoggy/grpc_course/hw_streaming_number_decomposition/NumberDecompositionpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
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

	service := NumberDecompositionpb.NewStreamGreetingServiceClient(conn)
	stream, err := service.NumberDecompositionStream(context.Background(), &
		NumberDecompositionpb.NumberDecompositionStreamRequest{Number: 120})
	if err != nil {
		logger.Fatal("error calling number decomposition stream")
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("got error from stream", zap.Error(err))
			break
		}

		logger.Info("got resp", zap.Any("resp", resp.String()))
	}
}
