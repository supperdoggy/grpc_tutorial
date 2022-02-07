package main

import (
	"context"
	"github.com/supperdoggy/grpc_course/streaming_greet/GreetStreampb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
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

	client := GreetStreampb.NewStreamGreetingServiceClient(conn)
	req := GreetStreampb.GreetManyTimesRequest{
		FirstName: "maksum",
		LastName: "marchyshak",
	}

	recv, err := client.GreetManyTimes(context.Background(), &req)
	if err != nil {
		logger.Fatal("error GreetManyTimes", zap.Error(err))
	}

	for {
		resp, err := recv.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Fatal("got fatal error", zap.Error(err))
		}
		logger.Info("got resp", zap.Any("resp", resp))
	}
}
