package test

import (
	"errors"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/net/pkg/core"
	"github.com/mlambda-net/net/pkg/local"
	"github.com/mlambda-net/net/pkg/remote"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Client_Future_Success(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }))
	s.Start(":9091")

	c := local.NewClient(":9091")
	sp := c.Spawn("dummy")
	r, e := sp.Future(&core.Response{Message: "1000"}, 5*time.Second).Result()
	d := r.(*core.Response)
	assert.Equal(t, "Good 1000", d.Message)
	assert.Nil(t, e)
	s.Stop()
}

func Test_Client_Future_Failure(t *testing.T) {
	s := remote.NewServer()

	s.Register("fail", actor.PropsFromFunc(func(ctx actor.Context) {
		ctx.Respond(errors.New("fail"))
	}))
	s.Start(":9002")
	c := local.NewClient(":9002")
	sp := c.Spawn("fail")
	r, e := sp.Future(&core.Response{Message: "Should fail"}, 5*time.Second).Result()
	assert.Nil(t, r)
	assert.Error(t, e)
	s.Stop()
}

func Test_Client_Future_Error(t *testing.T) {
	s := remote.NewServer()

	s.Register("fail", actor.PropsFromFunc(func(ctx actor.Context) {
		ctx.Respond(errors.New("this is a failure"))
	}))

	s.Start(":9003")
	c := local.NewClient(":9003")
	sp := c.Spawn("fail")
	r, e := sp.Future(&core.Response{Message: "Should fail"}, 5*time.Second).Result()
	assert.Nil(t, r)
	assert.Error(t, e)
	s.Stop()
}

func Test_Client_Send_Success(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }))
	s.Start(":9004")
	c := local.NewClient(":9004")
	sp := c.Spawn("dummy")
	sp.Send(&core.Response{Message: "1000"})
	s.Stop()
}

func Test_Client_Reconnecting(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }))
	s.Start(":9006")
	c := local.NewClient(":9006")
	sp := c.Spawn("dummy")
	sp.Send(&core.Response{Message: "1000"})
	time.Sleep(1000)
	s.Stop()
	s.Start(":9006")
	_, _ = sp.Future(&core.Response{Message: "1000"}, 5*time.Second).Result()
	s.Stop()
}
