package net

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/net/pkg/core"
	"github.com/mlambda-net/net/pkg/local"
	"time"
)

type Request interface {
	Token(token string) Request
	Request(message proto.Message) monad.Mono
	Send(message proto.Message)
}
type request struct {
	link  local.Address
	token string
}

func (r request) Token(token string) Request {
	r.token = token
	return r
}

type Client interface {
	Actor(name string) Request
}

type client struct {
	client local.Client
}

func (c client) Actor(name string) Request {
	return &request{ link: c.client.Spawn(name)}
}

func NewClient(remote string, port string) Client {
	l := local.NewClient(fmt.Sprintf("%s:%s", remote, port))
	return &client{
		client: l,
	}
}

func (r request) Request(data proto.Message) monad.Mono {
	f, e := r.link.Future(data, 4*time.Minute,r.token).Result()
	if e != nil {
		return monad.ToMono(e)
	}

	switch msg := f.(type) {
	case *core.Error:
		return monad.ToMono(fmt.Errorf("error: %d, message: %s", msg.Status, msg.Message))
	default:
		return monad.ToMono(msg)
	}
}

func (r request) Send(data proto.Message) {
	r.link.Send(data,r.token)
}
