package grpc

import (
	"github.com/sparrow-community/plugins/v4/config/source/grpc/proto"
	"go-micro.dev/v4/config/source"
)

type watcher struct {
	stream proto.Source_WatchService
}

func (w watcher) Next() (*source.ChangeSet, error) {
	wr, err := w.stream.Recv()
	if err != nil {
		return nil, err
	}
	return toChangeSet(wr.ChangeSet), err
}

func (w watcher) Stop() error {
	return w.stream.CloseSend()
}

func newWatcher(stream proto.Source_WatchService) *watcher {
	return &watcher{
		stream: stream,
	}
}
