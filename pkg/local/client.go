package local

import (
	"github.com/mlambda-net/net/pkg/common"
)

type Client interface {
	Spawn(s string) Address
}

type client struct {
	serializer common.Serializer
	server     string
}

func (c client) Spawn(kind string) Address {
	a := &address{
		kind:       kind,
		server:     c.server,
		serializer: c.serializer,
	}
	a.tryConnect()
	return a
}

func NewClient(server string) Client {
	return client{server: server, serializer: common.NewSerializer()}
}
