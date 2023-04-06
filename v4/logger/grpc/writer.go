package grpc

import (
	"context"
	"github.com/sparrow-community/plugins/v4/logger/grpc/proto"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"sync"
	"sync/atomic"
	"syscall"
)

type ZapGrpcWriter struct {
	serviceName string
	client      proto.LoggerService

	closed     int32
	closeMutex sync.Mutex
}

func (g *ZapGrpcWriter) Write(msg []byte) (n int, err error) {
	rsp, err := g.client.Write(context.Background(), &proto.WriteRequest{
		ServiceName: g.serviceName,
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

// InitializeLogger is initialize logger service
func InitializeLogger(serviceName string) {
	cs := client.DefaultClient
	l, err := NewLogger(
		WithServiceNameKey(serviceName),
		WithClientKey(proto.NewLoggerService("github.com.sparrow-community.logger-service", cs)),
	)
	if nil != err {
		logger.Error("logger service error: ", err)
	} else {
		logger.DefaultLogger = l
	}
}
