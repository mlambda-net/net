package test

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/net/pkg/core"
)

type dummy struct {

}

func (d *dummy) Receive(context actor.Context)  () {
  switch msg := context.Message().(type) {
  case *core.Response:
    context.Respond(&core.Response{
      Status:  200,
      Payload: msg.Payload,
      Message: "Good " + msg.Message,
    })

  case *core.Done:
    context.Respond(nil)

  case *core.Check:
    context.Respond(core.Unit())
  }
}
