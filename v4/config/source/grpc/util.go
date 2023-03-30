package grpc

import (
	"github.com/sparrow-community/plugins/v4/config/source/grpc/proto"
	"go-micro.dev/v4/config/source"
	"time"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}

func fromChangeSet(c *source.ChangeSet) *proto.ChangeSet {
	return &proto.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: c.Timestamp.Unix(),
		Source:    c.Source,
	}
}
