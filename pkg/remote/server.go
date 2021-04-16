package remote

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/net/pkg/common"
	"github.com/mlambda-net/net/pkg/core"
	"google.golang.org/grpc"
  "google.golang.org/grpc/health/grpc_health_v1"
  "log"
	"net"
)

type Server interface {
	Start(address string)
	Stop()
	Register(kind string, producer *actor.Props, isAuthenticate bool, roles []string)
	Check(status ...func(*Status) )
}

type server struct {
	props    map[string]*actor.Props
	srv      *grpc.Server
	lis      net.Listener
	register bool
	secure   map[string]*secure
	status   []func(*Status)
}

func (s *server) Check(status ...func(*Status)) {
	s.status = status
}

func (s *server) Register(kind string,producer *actor.Props, isAuthenticate bool, roles []string ) {
	s.props[kind] = producer
	s.secure[kind] = &secure{
		isAuth: isAuthenticate,
		roles: roles,
	}
}

func (s *server) Start(address string) {
  go func() {
    lis, err := net.Listen("tcp", address)
    if err != nil {
      log.Fatalf("failed to listen: %v", err)
    } else {
      s.lis = lis
      s.srv = grpc.NewServer()

      if !s.register {
        core.RegisterConnectorServer(s.srv, &service{
          system:    actor.NewActorSystem(),
          status:    s.status,
          props:     s.props,
          secure:    s.secure,
          pids:      make(map[string]*actor.PID),
          serialize: common.NewSerializer()})

        healthService := NewHealthChecker()
        grpc_health_v1.RegisterHealthServer(s.srv, healthService)

        s.register = true
      }
      if err := s.srv.Serve(s.lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
      }
    }
  }()
}

func (s *server) Stop() {
	if s.srv != nil {
		s.srv.Stop()
	}
	if s.lis != nil {
		_ = s.lis.Close()
	}

}

func NewServer() Server {
	return &server{register: false, props: make(map[string]*actor.Props), secure: make(map[string]*secure)}
}
