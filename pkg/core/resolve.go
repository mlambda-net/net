package core

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/ex"
  "github.com/sirupsen/logrus"
)

type Resolve interface {
  Mono (m monad.Mono, onSuccess  func(a  types.Any) interface{} )
}

type resolve struct {
  ctx actor.Context
}

func (r *resolve) Mono (m monad.Mono, onSuccess  func(a  types.Any) interface{} ) {
  p := uuid.New()
  resp, err := m.Unwrap()
  if err != nil {
    switch _err := err.(type) {
    case ex.Friendly:
      logrus.Info(_err.Fail().Error(), p.String())
    case ex.Crashed:
      logrus.Error(_err.Fail().Error(), p.String())
    default:
      logrus.Error(_err.Error(), p.String())
    }

    r.ctx.Respond(&Error{
      Message: err.Error(),
      Trace: p.String(),
    })
  } else {
    r.ctx.Respond(onSuccess(resp))
  }
}

func NewResolve(ctx actor.Context) Resolve  {
  return &resolve{ctx: ctx}
}
