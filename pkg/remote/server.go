package remote

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/net/pkg/common"
	"github.com/mlambda-net/net/pkg/core"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server interface {
	Start(address string)
	Stop()
	Register(kind string, producer *actor.Props)
}

type server struct {
	props    map[string]*actor.Props
	srv      *grpc.Server
	lis      net.Listener
	register bool
}

func (s *server) Register(kind string, producer *actor.Props) {
	s.props[kind] = producer
}

func (s *server) Start(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		s.lis = lis
		s.srv = grpc.NewServer()
		go func() {

			if !s.register {
				core.RegisterConnectorServer(s.srv, &service{ctx: actor.EmptyRootContext, props: s.props, pids: make(map[string]*actor.PID), serialize: common.NewSerializer()})
				s.register = true
			}
			if err := s.srv.Serve(s.lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()
	}
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
	return &server{register: false, props: make(map[string]*actor.Props)}
}
