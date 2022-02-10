package service

import (
	"github.com/supperdoggy/grpc_course/hw_client_streaming/avrgpb"
	"go.uber.org/zap"
	"io"
)

type IService interface {
	AvrgClientNumberStream(avrgpb.StreamGreetingService_AvrgClientNumberStreamServer) error
}

type service struct {
	l *zap.Logger
}

func NewService(l *zap.Logger) IService {
	return &service{l: l}
}

func (s *service) AvrgClientNumberStream(stream avrgpb.StreamGreetingService_AvrgClientNumberStreamServer) error {
	var (
		count int64
		sum int64
	)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&avrgpb.AvrgClientNumberStreamResponse{Number: sum/count})
		}

		if err != nil {
			return err
		}

		count++
		sum += req.GetNumber()
	}
}

