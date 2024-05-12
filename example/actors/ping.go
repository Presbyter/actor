package actors

import (
	"fmt"
	"log"

	"github.com/Presbyter/actor"
)

type pingActor struct {
	mailbox chan any
	state   struct {
		name string
	}
}

func (a *pingActor) Name() string {
	if a.state.name == "" {
		a.state.name = "ping"
	}
	return a.state.name
}

func (a *pingActor) Send(msg any) {
	if a.mailbox == nil {
		a.mailbox = make(chan any)
	}
	a.mailbox <- msg
}

func (a *pingActor) Start() {
	for msg := range a.mailbox {
		fmt.Println("PING. Send msg:", msg)
		s, ok := actor.GetSystem("PingPong")
		if !ok {
			log.Fatal("PingPong actor system not found")
		}
		pong := s.GetActor("user1/pong")
		if pong == nil {
			log.Fatal("user1/ping actor not found")
		}
		pong.Send(msg)
	}
}

func NewPingActor(name string) actor.ActorRef {
	return &pingActor{
		mailbox: make(chan any),
		state:   struct{ name string }{name: name},
	}
}
