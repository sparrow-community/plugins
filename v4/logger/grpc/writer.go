package grpc

import (
	"context"
	"github.com/sparrow-community/plugins/v4/logger/grpc/proto"
	"sync"
	"sync/atomic"
	"syscall"
)

type ZapGrpcWriter struct {
	ServiceName string
	Client      proto.LoggerService

	closed     int32
	closeMutex sync.Mutex
}

func (g *ZapGrpcWriter) Write(msg []byte) (n int, err error) {
	rsp, err := g.Client.Write(context.Background(), &proto.WriteRequest{
		ServiceName: g.ServiceName,
		Data:        msg,
	})
	if err != nil {
		return 0, err
	}
	return int(rsp.N), nil
}

func (g *ZapGrpcWriter) Sync() error {
	return nil
}

func (g *ZapGrpcWriter) Close() error {
	g.closeMutex.Lock()
	defer g.closeMutex.Unlock()

	if g.Closed() {
		return syscall.EINVAL
	}

	atomic.StoreInt32(&g.closed, 1)
	return nil
}

func (g *ZapGrpcWriter) Closed() bool {
	return atomic.LoadInt32(&g.closed) != 0
}
