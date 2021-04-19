package core

import (
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/ex"
  "github.com/sirupsen/logrus"
)

type Resolve interface {
  Mono (m monad.Mono) Resolver
  Error(err error) Resolver
}

type resolve struct {
}

type Resolver interface {
  Then(apply func(types.Any) types.Any) Resolver
  Response() interface{}
}

type resolver struct {
  response types.Any
  bind     func(types.Any) types.Any
}

func (r *resolver) Then(apply func(types.Any) types.Any) Resolver {
  r.bind = apply
  return r
}

func (r *resolver) Response() interface{} {
  switch mr := r.response.(type) {
  case *Error:
    return mr
  default:
    return r.bind(r.response)
  }
}

func (r *resolve) Error(err error) Resolver {
  p := uuid.New()
  r.handleError(err, p.String())
  return &resolver{
    response: &Error{
      Message: err.Error(),
      Trace:   p.String(),
    },
  }
}

func (r *resolve) Mono (m monad.Mono) Resolver {
  p := uuid.New()
  resp, err := m.Unwrap()
  if err != nil {
    r.handleError(err, p.String())
    return &resolver{
      response: &Error{
        Message: err.Error(),
        Trace:   p.String(),
      },
    }
  }
  return &resolver{
    response: resp,
  }
}

func (r *resolve) handleError(err error, trace string) {
  switch _err := err.(type) {
  case ex.Friendly:
    logrus.WithFields(logrus.Fields{
      "trace": trace,
    }).Infoln(_err.Fail().Error(), trace)
  case ex.Crashed:
    logrus.WithFields(logrus.Fields{
      "trace": trace,
    }).Errorln(_err.Fail().Error())
  case ex.Exception:
    logrus.WithFields(logrus.Fields{
      "trace": trace,
    }).Warningln(_err.Error())
  default:
    logrus.WithFields(logrus.Fields{
      "trace": trace,
    }).Error(_err.Error())
  }
}

func NewResolve() Resolve  {
  return &resolve{}
}
