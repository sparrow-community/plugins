package grpc

import (
	context "context"
	"github.com/sparrow-community/plugins/v4/logger/grpc/proto"
	"go-micro.dev/v4/client"
)

type grpcLogger struct {
	client proto.LoggerService
}

func (g grpcLogger) Write(ctx context.Context, opts ...client.CallOption) (proto.Logger_WriteService, error) {
	//TODO implement me
	panic("implement me")
}
