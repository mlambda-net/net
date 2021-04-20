package core

import (
  "errors"
  "fmt"
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/ex"
  "github.com/sirupsen/logrus"
)

type Resolve interface {
  Mono (m monad.Mono) Resolver
  Monos(ms []monad.Mono) Resolver
  Error(err error) Resolver
  Maybe(m monad.Maybe) Resolver
  Maybes(ms []monad.Maybe) Resolver
  Fail(msg string) Resolver
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

func (r *resolve) Fail(msg string) Resolver  {
  trace := uuid.New()
  e := errors.New(msg)
  r.handleError(e, trace.String())
  return &resolver{
    response: e,
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

func (r *resolve) Monos(ms []monad.Mono) Resolver  {
  values := make([]interface{}, 0)
  errors := make([]error, 0)
  for _, m := range ms{
    res, e := m.Unwrap()
    if e != nil {
      errors = append(errors, e)
    } else {
      values = append(values, res)
    }
  }

  if len(errors) > 0 {
    trace := uuid.New()
    msg := fmt.Sprintf("%v", errors)

    logrus.WithFields(logrus.Fields{
      "trace": trace,
    }).Errorln(msg)

    return &resolver{
      response: &Error{
        Message: msg,
        Trace:  trace.String(),
      },
    }
  } else {
    return &resolver{
      response: values,
    }
  }
}

func (r *resolve) Maybes(ms []monad.Maybe) Resolver {
  values := make([]interface{}, 0)
  for _, m := range ms {
    switch ma := m.(type) {
    case monad.Just:
      values = append(values, ma.Value())
    }
  }

  if len(values) == 0 {
    return &resolver{
      response: &Empty{},
    }
  } else {
    return &resolver{response: values}
  }
}

func (r *resolve) Maybe(m monad.Maybe) Resolver {
  switch v := m.(type) {
  case monad.Just:
    return &resolver{
      response: v.Value(),
    }
  default:
    return &resolver{
      response: &Empty{},
    }
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
