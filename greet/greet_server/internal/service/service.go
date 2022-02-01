package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/supperdoggy/grpc_course/greet/greetpb"
	"go.uber.org/zap"
)

type server struct {
	l *zap.Logger
}

func NewServer(l *zap.Logger) greetpb.GreetServiceServer {
	return server{ l:l}
}

func (s server) SendGreeting(ctx context.Context, in *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	result := greetpb.GreetingResponse{}
	if in.Greeting.GetLastName() == "" || in.Greeting.GetFirstName() == "" {
		s.l.Error("got empty first_name or last_name", zap.Any("greeting", in.Greeting))
		return nil, errors.New("you need to fill both fisrt_name and last_name")
	}
	result.Result = fmt.Sprintf("Hello %s %s", in.Greeting.GetFirstName(), in.Greeting.GetLastName())
	return &result, nil
}
