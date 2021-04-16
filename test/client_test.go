package test

import (
	"errors"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/net/pkg/core"
	"github.com/mlambda-net/net/pkg/local"
	"github.com/mlambda-net/net/pkg/net"
	"github.com/mlambda-net/net/pkg/remote"
	"github.com/mlambda-net/net/pkg/security"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func Test_Client_Future_Success(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),true, nil)
	s.Start(":9091")

	c := net.NewClient("localhost", "9091")
	tk := security.NewToken(os.Getenv("SECRET_KEY"))
	token, err := tk.Create(map[string]interface{}{"authorize": true})
	assert.Nil(t, err)

	dummy := c.Actor("dummy")
	r, e := dummy.Token(fmt.Sprintf("bearer %s",token)).Request(&core.Response{Message: "1000"}).Unwrap()


	d := r.(*core.Response)
	assert.Equal(t, "Good 1000", d.Message)
	assert.Nil(t, e)
	s.Stop()
}

func Test_Client_Future_Failure(t *testing.T) {
	s := remote.NewServer()

	s.Register("fail", actor.PropsFromFunc(func(ctx actor.Context) {
		ctx.Respond(errors.New("fail"))
	}), false, []string{""})
	s.Start(":9002")
	c := local.NewClient(":9002")
	sp := c.Spawn("fail")
	r, e := sp.Future(&core.Response{Message: "Should fail"}, 5*time.Second,"").Result()
	assert.Nil(t, r)
	assert.Error(t, e)
	s.Stop()
}

func Test_Client_Future_Error(t *testing.T) {
	s := remote.NewServer()

	s.Register("fail", actor.PropsFromFunc(func(ctx actor.Context) {
		ctx.Respond(errors.New("this is a failure"))
	}), false, []string{""})

	s.Start(":9003")
	c := local.NewClient(":9003")
	sp := c.Spawn("fail")
	r, e := sp.Future(&core.Response{Message: "Should fail"}, 5*time.Second, "").Result()
	assert.Nil(t, r)
	assert.Error(t, e)
	s.Stop()
}

func Test_Client_Send_Success(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),false, []string{""})
	s.Start(":9004")
	c := local.NewClient(":9004")
	sp := c.Spawn("dummy")
	sp.Send(&core.Response{Message: "1000"},"")
	s.Stop()
}

func Test_Client_Reconnecting(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }), false, []string{""})
	s.Start(":9006")
	c := local.NewClient(":9006")
	sp := c.Spawn("dummy")
	sp.Send(&core.Response{Message: "1000"},"")
	time.Sleep(1000)
	s.Stop()
	s.Start(":9006")
	_, _ = sp.Future(&core.Response{Message: "1000"}, 5*time.Second, "").Result()
	s.Stop()
}


func Test_Client_Future_Auth_Fail(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),true, []string{""})
	s.Start(":9091")

	c := local.NewClient(":9091")
	sp := c.Spawn("dummy")
	_, e := sp.Future(&core.Response{Message: "1000"}, 5*time.Second, "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU2MzgwMjAsInVzZXJfaWQiOiJjb3lvdGVAYWNtZS5jb20ifQ.wLQX8pmzqqxQU5RWxRo6hbQmLfQeiKQs9iFhx5p3Czg").Result()
	assert.Equal(t, "Token is expired", e.Error())
	s.Stop()
}


func Test_Client_Health_And_Live(t *testing.T) {
	s := remote.NewServer()
	s.Check(func( status *remote.Status )  {
		status.Add(true, "check func", "ok")
	})
	s.Start(":9091")
	c := local.NewClient(":9091")
	status, err := c.Live()
	assert.Nil(t, err)
	assert.Equal(t, "200 ok", status.Message)

	status, err = c.Health()
	assert.Nil(t, err)
	assert.Equal(t, "check func is ok", status.Message)
}


func Test_Client_Success(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),false, []string{""})
	s.Check(func( status *remote.Status )  {
		status.Add(true, "check func", "ok")
	})
	s.Start(":9093")

	c := net.NewClient("localhost", "9093")
	r, e := c.Actor("dummy").Request(&core.Response{ Message: "Great"}).Unwrap()
	assert.Nil(t, e)
	assert.Equal(t, "Good Great", r.(*core.Response).Message)
}

func Test_Client_Fail(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),true, []string{""})
	s.Check(func( status *remote.Status )  {
		status.Add(true, "check func", "ok")
	})
	s.Start(":9092")

	c := net.NewClient("localhost", "9092")
	r, e := c.Actor("dummy").Token("bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU2MzgwMjAsInVzZXJfaWQiOiJjb3lvdGVAYWNtZS5jb20ifQ.wLQX8pmzqqxQU5RWxRo6hbQmLfQeiKQs9iFhx5p3Czg").Request(&core.Response{ Message: "Great"}).Unwrap()
	assert.Nil(t, r)
	assert.Equal(t, "Token is expired", e.Error())
}

func Test_Client_Send_Done(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),false, []string{""})
	s.Start(":9010")
	c := local.NewClient(":9010")
	sp := c.Spawn("dummy")
	sp.Send(&core.Done{},"")
	s.Stop()
}

func Test_Client_Send_Check(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),false, []string{""})
	s.Start(":9010")
	c := local.NewClient(":9010")
	sp := c.Spawn("dummy")
	sp.Send(&core.Check{},"")
	s.Stop()
}

func Test_Client_Check(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),false, []string{""})
	s.Start(":9011")

	c := net.NewClient("localhost", "9011")
	status, err := c.Live()
	assert.Nil(t, err)
	assert.Equal(t, true, status.Success)
	s.Stop()
}

func Test_Client_Health(t *testing.T) {
	s := remote.NewServer()
	s.Register("dummy", actor.PropsFromProducer(func() actor.Actor { return &dummy{} }),false, []string{""})
	s.Start(":9011")

	c := net.NewClient("localhost", "9011")
	status, err := c.Health()
	assert.Nil(t, err)
	assert.Equal(t, true, status.Success)
	s.Stop()
}
