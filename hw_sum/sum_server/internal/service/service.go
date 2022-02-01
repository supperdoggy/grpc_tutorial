package service

import (
	"context"
	"github.com/supperdoggy/grpc_course/hw_sum/sumpb"
	"go.uber.org/zap"
)

type IService interface {
	Add(ctx context.Context, in *sumpb.AddRequest) (*sumpb.AddResponse, error)
}

type service struct {
	l *zap.Logger
}

func NewService(l *zap.Logger) IService {
	return service{l: l}
}

func (s service) Add(ctx context.Context, in *sumpb.AddRequest) (*sumpb.AddResponse, error) {
	result := sumpb.AddResponse{Sum: in.GetA() + in.GetB()}
	return &result, nil
}
