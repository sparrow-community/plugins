package grpc

import (
	"context"
	"github.com/sparrow-community/plugins/v4/logger/grpc/proto"
	"go-micro.dev/v4/logger"
)

type Writer struct {
	serviceName string
	client      proto.LoggerService
	message     chan []byte
}

func (g *Writer) Write(msg []byte) error {
	go func() {
		g.message <- msg
	}()
	return nil
}

func (g *Writer) write() error {
	stream, err := g.client.Write(context.Background())
	if err != nil {
		return err
	}
	defer stream.Close()
	defer close(g.message)
	go func() {
		for {
			select {
			case msg := <-g.message:
				err := stream.Send(&proto.WriteRequest{
					ServiceName: g.serviceName,
					Data:        msg,
				})
				if err != nil {
					logger.Errorf("failed to send message: %v", err)
				}
			}
		}
	}()
	return nil
}
