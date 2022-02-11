package service

import (
	"github.com/supperdoggy/grpc_course/hw_bidirectional_streaming/maxpb"
	"go.uber.org/zap"
	"io"
)

type IService interface {
	MaxStreamInt(maxpb.BiDirectionalStreaming_MaxStreamIntServer) error
}

type service struct {
	l *zap.Logger
}

func NewService(l *zap.Logger) IService {
	return &service{l: l}
}

func (s *service) MaxStreamInt(stream maxpb.BiDirectionalStreaming_MaxStreamIntServer) error {
	var max *int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if max == nil {
			a := req.GetNumber()
			max = &a
			err = stream.Send(&maxpb.MaxStreamIntResponse{Number: *max})
			if err != nil {
				return err
			}
			s.l.Info("sent max", zap.Any("max", max))
			continue
		}

		if req.GetNumber() > *max {
			*max = req.GetNumber()
			err = stream.Send(&maxpb.MaxStreamIntResponse{Number: *max})
			if err != nil {
				return err
			}
			s.l.Info("sent max", zap.Any("max", max))
		}
	}
}

