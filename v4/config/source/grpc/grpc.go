package grpc

import (
	"github.com/sparrow-community/protos/config"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/logger"
	"golang.org/x/net/context"
	"os"
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

	var sourceService proto.SourceService
	path := "/"
	if options.Context != nil {
		c, ok := options.Context.Value(clientKey{}).(proto.SourceService)
		if ok {
			sourceService = c
		}
		p, ok := options.Context.Value(pathKey{}).(string)
		if ok {
			path = p
		}
	}

	return &grpcSource{
		opts:   options,
		path:   path,
		client: sourceService,
	}
}

func InitializeConfig(path string, client client.Client, sources ...source.Source) (config.Config, error) {
	var _sources []source.Source
	// register center
	cfgClient := proto.NewSourceService("github.com.sparrow-community.config-service", client)
	registerCenterSource := NewSource(
		WithPath(path),
		WithClient(cfgClient),
	)
	_sources = append(_sources, registerCenterSource)

	// local
	b, err := exists(path)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if b {
		_sources = append(_sources, file.NewSource(file.WithPath(path)))
	}

	if len(sources) != 0 {
		_sources = append(_sources, sources...)
	}

	if err := config.Load(_sources...); err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return config.DefaultConfig, err
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
