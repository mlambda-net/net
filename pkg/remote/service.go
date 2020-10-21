package remote

import (
  "context"
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/net/pkg/common"
  "reflect"
  "strings"
  "time"
)
import "github.com/mlambda-net/net/pkg/core"
type service struct {
  core.UnimplementedConnectorServer
  serialize common.Serializer
  props map[string]*actor.Props
  ctx   *actor.RootContext
  pids map[string]*actor.PID
}



func (s *service) Call(c context.Context, r *core.Request) (*core.Response, error) {
  message, err := s.serialize.Deserialize(r.Type, r.Payload)
  if err != nil {
    return nil, err
  }

  if _, ok := s.pids[r.Kind]; !ok {
    if prop, ok := s.props[r.Kind]; ok {
      pid := s.ctx.Spawn(prop)
      s.pids[r.Kind] = pid
    }
  }

  if pid, ok := s.pids[r.Kind]; ok {
    v, e := s.ctx.RequestFuture(pid, message, 10*time.Second).Result()
    if e != nil {
      return nil, e
    }

    data, err := s.serialize.Serialize(v)

    if err != nil {
      return nil, err
    }

    return &core.Response{
      Status:  200,
      Payload: data,
      Message: strings.ReplaceAll(reflect.TypeOf(v).String(), "*", ""),
    }, nil
  }

  return &core.Response{
    Status:  400,
    Message: "can not find the kind",
  }, nil

}


