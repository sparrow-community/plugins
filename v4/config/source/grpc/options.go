package grpc

import (
	"github.com/sparrow-community/plugins/v4/config/source/grpc/proto"
	"go-micro.dev/v4/config/source"
	"golang.org/x/net/context"
)

type clientKey struct{}
type pathKey struct{}

func WithClient(sourceService proto.SourceService) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, clientKey{}, sourceService)
	}
}

func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, pathKey{}, p)
	}
}
