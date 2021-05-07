package middelware

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func Auth (token string)  func (next actor.ReceiverFunc) actor.ReceiverFunc {
	return func(next actor.ReceiverFunc) actor.ReceiverFunc {
		return func(c actor.ReceiverContext, envelope *actor.MessageEnvelope) {
			envelope.Header = actor.EmptyMessageHeader
			envelope.Header.Set("token", token)
			next(c, envelope)
		}
	}
}
