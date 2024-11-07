package service

import (
	"context"
	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/demo/demo_proto/kitex_gen/pbapi"
	"testing"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	req := &pbapi.Request{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
