package service

import (
	"context"

	v1 "github.com/wizard-xie/realworld/api/realworld/v1"
	"github.com/wizard-xie/realworld/internal/biz"
)

// RealworldService is a greeter service.
type RealworldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealworldUsecase
}

// NewRealworldService new a greeter service.
func NewRealworldService(uc *biz.RealworldUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}

// SayHello implements realworld.GreeterServer.
func (s *RealworldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Realworld{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
