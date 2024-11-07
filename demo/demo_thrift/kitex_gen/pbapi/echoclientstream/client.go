// Code generated by Kitex v0.9.1. DO NOT EDIT.

package echoclientstream

import (
	"context"
	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/client/streamclient"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	EchoClientStream(ctx context.Context, callOptions ...callopt.Option) (stream EchoClientStream_EchoClientStreamClient, err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	EchoClientStream(ctx context.Context, callOptions ...streamcall.Option) (stream EchoClientStream_EchoClientStreamClient, err error)
}

type EchoClientStream_EchoClientStreamClient interface {
	streaming.Stream
	Send(*pbapi.Request) error
	CloseAndRecv() (*pbapi.Response, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClientStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoClientStreamClient struct {
	*kClient
}

func (p *kEchoClientStreamClient) EchoClientStream(ctx context.Context, callOptions ...callopt.Option) (stream EchoClientStream_EchoClientStreamClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoClientStream(ctx)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClientStreamStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoClientStreamStreamClient struct {
	*kClient
}

func (p *kEchoClientStreamStreamClient) EchoClientStream(ctx context.Context, callOptions ...streamcall.Option) (stream EchoClientStream_EchoClientStreamClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.EchoClientStream(ctx)
}
