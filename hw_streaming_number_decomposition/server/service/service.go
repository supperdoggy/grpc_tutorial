package service

import (
	"github.com/supperdoggy/grpc_course/hw_streaming_number_decomposition/NumberDecompositionpb"
	"go.uber.org/zap"
)

type IService interface {
	NumberDecompositionStream(*NumberDecompositionpb.NumberDecompositionStreamRequest, NumberDecompositionpb.StreamGreetingService_NumberDecompositionStreamServer) error
}

type service struct {
	l *zap.Logger
}

func NewService(l *zap.Logger) IService {
	return &service{l: l}
}

func (s *service) NumberDecompositionStream(req *NumberDecompositionpb.NumberDecompositionStreamRequest, stream NumberDecompositionpb.StreamGreetingService_NumberDecompositionStreamServer) error {
	k := int64(2)
	n := req.GetNumber()
	for n > 1 {
		if n % k == 0{
			err := stream.Send(&NumberDecompositionpb.NumberDecompositionStreamResponse{Number: k})
			if err != nil {
				return err
			}
			n = n / k
		} else {
			k++
		}
	}
	return nil
}