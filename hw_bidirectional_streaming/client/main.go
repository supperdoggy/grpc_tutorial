package main

import (
	"context"
	"github.com/supperdoggy/grpc_course/hw_bidirectional_streaming/maxpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"time"
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

	service := maxpb.NewBiDirectionalStreamingClient(conn)

	stream, err := service.MaxStreamInt(context.Background())
	if err != nil {
		logger.Fatal("error stream", zap.Error(err))
	}

	quit := make(chan struct{})
	go func() {
		for _, v := range []int64{1,5,3,6,2,20} {
			err := stream.Send(&maxpb.MaxStreamIntRequest{Number: v})
			if err != nil {
				logger.Fatal("error", zap.Error(err))
			}
			time.Sleep(2 * time.Second)
		}
	} ()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				quit <- struct{}{}
				return
			}

			if err != nil {
				logger.Fatal("got error", zap.Error(err))
			}

			logger.Info("got resp", zap.Any("response", resp.String()))
		}
	}()

	_ = <- quit
}
