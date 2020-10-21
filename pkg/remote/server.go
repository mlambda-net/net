package remote

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/net/pkg/common"
  "github.com/mlambda-net/net/pkg/core"
  "google.golang.org/grpc"
  "log"
  "net"
  "sync"
)

type Server interface {
  Start()
  Stop()
  Register(kind string, producer *actor.Props)
}

type server struct {
  wg  *sync.WaitGroup
  props map[string]*actor.Props
}

func (s *server) Register(kind string, producer *actor.Props) {
  s.props[kind] = producer
}

func (s server) Start() {

  go func() {
    lis, err := net.Listen("tcp", ":9001")
    if err != nil {
      log.Fatalf("failed to listen: %v", err)
      s.wg.Done()
    }
    srv := grpc.NewServer()
    core.RegisterConnectorServer(srv, &service{ctx: actor.EmptyRootContext, props: s.props, pids: make(map[string]*actor.PID), serialize: common.NewSerializer()})
    if err := srv.Serve(lis); err != nil {
      log.Fatalf("failed to serve: %v", err)
      s.wg.Done()
    }
    s.wg.Wait()
  }()
}

func (s server) Stop() {
  s.wg.Done()
}

func NewServer() Server  {
  wg := new(sync.WaitGroup)
  wg.Add(1)
  return &server{wg: wg, props: make(map[string]*actor.Props)}
}
