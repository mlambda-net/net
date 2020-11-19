package local

import (
	"context"
	"github.com/mlambda-net/net/pkg/common"
	"github.com/mlambda-net/net/pkg/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

type Client interface {
	Spawn(s string) Address
	Live()  (*core.Status, error)
	Health()  (*core.Status, error)
}



type client struct {
	serializer common.Serializer
	server     string
	client     core.ConnectorClient
	conn       *grpc.ClientConn
}


func (c client) Live() (*core.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	return c.client.Live(ctx, &core.Check{})
}

func (c client) Health()  (*core.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	return c.client.Health(ctx, &core.Check{})
}

func (c client) Spawn(kind string) Address {
	a := &address{
		kind:       kind,
		server:     c.server,
		serializer: c.serializer,
	}
	a.tryConnect(c.conn, c.client)
	return a
}

func NewClient(server string) Client {
	c := client{server: server, serializer: common.NewSerializer()}
	c.conn, c.client = createConnection(server)
	return c
}


func  createConnection(server string) (*grpc.ClientConn, core.ConnectorClient) {
	ka := keepalive.ClientParameters{
		Time:                20 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithKeepaliveParams(ka))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := core.NewConnectorClient(conn)
	return conn, client
}