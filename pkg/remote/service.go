package remote

import (
	"context"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/mlambda-net/net/pkg/common"
	"github.com/mlambda-net/net/pkg/remote/middelware"
	"github.com/mlambda-net/net/pkg/security"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)
import "github.com/mlambda-net/net/pkg/core"

type secure struct {

	isAuth bool
	roles []string
}

type result struct {
	name    interface{}
	success bool
	message string
}

type Status struct {
	results []result
}

func (s *Status) Add(success bool, name string, message string) {
	s.results = append(s.results, result{
		success: success,
		name: name,
		message: message,
	})
}

func (s *Status) getText() string {
	var sb strings.Builder
	for _, s := range s.results {
		if s.success {
			sb.WriteString(fmt.Sprintf("%s is ok ", s.name))
		} else {
			sb.WriteString(fmt.Sprintf("fail %s message: %s ", s.name, s.message))
		}
	}
	return sb.String()
}

func (s *Status) getStatus() bool {
	check := true
	for _, s:= range s.results {
		if !s.success {
			check = false
		}
	}
	return check

}

type service struct {
	core.UnimplementedConnectorServer
	serialize common.Serializer
	props     map[string]*actor.Props
	pids      map[string]*actor.PID
	secure    map[string]*secure
	status    []func(status *Status)
	system    *actor.ActorSystem
}

func (s *service) Call(_ context.Context, r *core.Request) (*core.Response, error) {

  defer func() {
    if err := recover(); err != nil {
      log.Println("panic occurred:", err)
    }
  }()

  message, err := s.serialize.Deserialize(r.Type, r.Payload)
  if err != nil {

    return &core.Response{
      Status:  500,
      Message: err.Error(),
    }, nil
  }

  if prop, ok := s.props[r.Kind]; ok {
    secure := s.secure[r.Kind]
    if secure.isAuth {
      identity, err := security.NewIdentity(r.Token)
      if err != nil {
        return &core.Response{
          Status:  http.StatusUnauthorized,
          Message: err.Error(),
        }, nil
      }

      if identity.Authenticate() {
        if identity.HasRoles(secure.roles) {
          headers := identity.Serialize()
          return s.exec(s.system.Root.Spawn(prop.WithReceiverMiddleware(middelware.Auth(headers))), message)
        } else {
          return &core.Response{
            Status:  http.StatusUnauthorized,
            Message: "The user don't have access",
          }, nil
        }
      }
      return &core.Response{
        Status:  http.StatusUnauthorized,
        Message: "User is not authenticated",
      }, nil
    } else {
      prop := s.props[r.Kind].WithReceiverMiddleware(middelware.Auth(""))
      return s.exec(s.system.Root.Spawn(prop), message)
    }
  }

  return &core.Response{
    Status:  http.StatusNotFound,
    Message: "there is not actor with that kind",
  }, nil
}

func (s *service) exec(pid *actor.PID, message interface{}) (*core.Response, error) {
	v, e := s.system.Root.RequestFuture(pid, message, 10*time.Second).Result()
	if e != nil {
		return &core.Response{
			Status:  http.StatusInternalServerError,
			Message: e.Error(),
		}, nil
	}

	if err, ok := v.(error); ok {
		return &core.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	if v == nil {
		v = &core.Done{}
	}

	data, err := s.serialize.Serialize(v)
	if err != nil {
		return &core.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}
	return &core.Response{
		Status:  http.StatusOK,
		Payload: data,
		Message: strings.ReplaceAll(reflect.TypeOf(v).String(), "*", ""),
	}, nil
}

func (s* service) Live(_ context.Context, _ *core.Check) (*core.Status, error) {
	return &core.Status{Success: true, Message: fmt.Sprintf("%d ok", http.StatusOK)}, nil
}

func (s* service) Health(_ context.Context, _ *core.Check) (*core.Status, error)  {
	status := &Status{}
	for _, f := range s.status {
		f(status)
	}

	return &core.Status{Success: status.getStatus(), Message: status.getText()}, nil
}
