package local

import (
	"context"
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/net/pkg/common"
	"github.com/mlambda-net/net/pkg/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"reflect"
	"strings"
	"time"
)

type Address interface {
	Future(proto.Message, time.Duration) monad.Future
	Send(proto.Message)
}

type address struct {
	client     core.ConnectorClient
	serializer common.Serializer
	kind       string
	server     string
	conn       *grpc.ClientConn
}

func (a *address) Future(message proto.Message, timeout time.Duration) monad.Future {

	f := monad.NewFuture()
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		data, err := a.serializer.Serialize(message)
		if err != nil {
			f.SetResult(err)
		}

		t := reflect.TypeOf(message)

		res, err := a.client.Call(ctx, &core.Request{
			Type:    strings.ReplaceAll(t.String(), "*", ""),
			Payload: data,
			Kind:    a.kind,
		})

		if err != nil {
			f.SetResult(err)
		} else {
			if res.Status != 200 {
				f.SetResult(errors.New(res.Message))
			} else {
				v, e := a.serializer.Deserialize(res.Message, res.Payload)
				if e != nil {
					f.SetResult(err)
				} else {
					f.SetResult(v)
				}
			}
		}
	}()
	return f
}

func (a *address) Send(message proto.Message) {
	_, e := a.Future(message, 5*time.Second).Result()
	if e != nil {
		log.Fatal(e)
	}
}

func (a *address) tryConnect() {
	if a.conn == nil {
		conn, client := a.createConnection()
		a.conn = conn
		a.client = client
	}
}

func (a *address) createConnection() (*grpc.ClientConn, core.ConnectorClient) {
	ka := keepalive.ClientParameters{
		Time:                20 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}
	conn, err := grpc.Dial(a.server, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithKeepaliveParams(ka))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := core.NewConnectorClient(conn)
	return conn, client
}
