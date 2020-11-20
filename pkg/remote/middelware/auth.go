package middelware

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func Auth (headers map[string]string)  func (next actor.ReceiverFunc) actor.ReceiverFunc {
	return func(next actor.ReceiverFunc) actor.ReceiverFunc {
		return func(c actor.ReceiverContext, envelope *actor.MessageEnvelope) {
			envelope.Header = actor.EmptyMessageHeader
			for k, v := range headers {
				envelope.Header.Set(k, v)
			}
			next(c, envelope)
		}
	}
}
