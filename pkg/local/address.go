package local

import (
  "context"
  "errors"
  "github.com/gogo/protobuf/proto"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/common"
  "github.com/mlambda-net/net/pkg/core"
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
}


func (a address) Future(message proto.Message,timeout time.Duration) monad.Future {

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
    }

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
  }()
  return f
}

func (a address) Send(message proto.Message) {
  _, e:=  a.Future(message, 5*time.Second).Result()
  if e != nil {
    log.Fatal(e)
  }
}
