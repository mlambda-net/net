package local

import (
	"github.com/mlambda-net/net/pkg/common"
	"github.com/mlambda-net/net/pkg/core"
	"google.golang.org/grpc"

	"log"
)

type Client interface {
	Spawn(s string) Address
}

type client struct {
	client     core.ConnectorClient
	serializer common.Serializer
}

func (c client) Spawn(kind string) Address {
	return address{
		kind:       kind,
		client:     c.client,
		serializer: c.serializer,
	}
}

func NewClient(server string) Client {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := core.NewConnectorClient(conn)
	return client{client: c, serializer: common.NewSerializer()}
}
