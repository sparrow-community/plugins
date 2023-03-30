package grpc

import (
	"github.com/sparrow-community/plugins/v4/config/source/grpc/proto"
	"go-micro.dev/v4/config/source"
	"golang.org/x/net/context"
)

type grpcSource struct {
	opts   source.Options
	client proto.SourceService
	path   string
}

func (g *grpcSource) Read() (*source.ChangeSet, error) {
	rds, err := g.client.Read(context.Background(), &proto.ReadRequest{
		Path: g.path,
	})
	if err != nil {
		return nil, err
	}
	return toChangeSet(rds.ChangeSet), err
}

func (g *grpcSource) Write(set *source.ChangeSet) error {
	_, err := g.client.Write(context.Background(), &proto.WriteRequest{
		ChangeSet: fromChangeSet(set),
	})
	return err
}

func (g *grpcSource) Watch() (source.Watcher, error) {
	wds, err := g.client.Watch(context.Background(), &proto.WatchRequest{
		Path: g.path,
	})
	if err != nil {
		return nil, err
	}

	return newWatcher(wds), nil
}

func (g *grpcSource) String() string {
	return "grpc"
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	var client proto.SourceService
	path := "/"
	if options.Context != nil {
		c, ok := options.Context.Value(clientKey{}).(proto.SourceService)
		if ok {
			client = c
		}
		p, ok := options.Context.Value(pathKey{}).(string)
		if ok {
			path = p
		}
	}

	return &grpcSource{
		opts:   options,
		path:   path,
		client: client,
	}
}
