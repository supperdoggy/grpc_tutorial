package main

import (
	"context"
	"github.com/supperdoggy/grpc_course/greet/greetpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewDevelopment()

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("error dialing with server", zap.Error(err))
	}

	c := greetpb.NewGreetServiceClient(cc)

	result, err := c.SendGreeting(context.Background(), &greetpb.GreetingRequest{Greeting: &greetpb.Greeting{FirstName: "Maksym", LastName: "Marchyshak"}})
	logger.Info("got result", zap.Any("result", result.String()), zap.Error(err))
	defer cc.Close()

}
