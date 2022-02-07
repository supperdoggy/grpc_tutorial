package service

import (
	"fmt"
	"github.com/supperdoggy/grpc_course/streaming_greet/GreetStreampb"
	"go.uber.org/zap"
	"time"
)

type IServer interface {
	GreetManyTimes(*GreetStreampb.GreetManyTimesRequest, GreetStreampb.StreamGreetingService_GreetManyTimesServer) error
}

type server struct {
	l *zap.Logger
}

func NewServer(l *zap.Logger) IServer {
	return server{l: l}
}

func (s server) GreetManyTimes(req *GreetStreampb.GreetManyTimesRequest, stream GreetStreampb.StreamGreetingService_GreetManyTimesServer) error {
	result := fmt.Sprintf("Hello %s %s", req.GetFirstName(), req.GetLastName())
	for k := 0; k < 10; k++{
		resp := GreetStreampb.GreetManyTimesResponse{Answer: result}
		err := stream.Send(&resp)
		if err != nil {
			s.l.Error("error sending response", zap.Error(err))
			return err
		}
		time.Sleep(2*time.Second)
	}
	return nil
}
