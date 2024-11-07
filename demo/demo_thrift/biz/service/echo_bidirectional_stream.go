package service

import (
	"context"
	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
)

type EchoBidirectionalStreamService struct {
	ctx context.Context
}

// NewEchoBidirectionalStreamService new EchoBidirectionalStreamService
func NewEchoBidirectionalStreamService(ctx context.Context) *EchoBidirectionalStreamService {
	return &EchoBidirectionalStreamService{ctx: ctx}
}

func (s *EchoBidirectionalStreamService) Run(stream pbapi.EchoBidirectionalStream_EchoBidirectionalStreamServer) (err error) {
	return
}
