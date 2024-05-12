package actors

import (
	"fmt"
	"log"

	"github.com/Presbyter/actor"
)

type pingActor struct {
	actor.BaseActor
	mailbox chan string
}

func (a *pingActor) Send(msg any) {
	v, ok := msg.(string)
	if !ok {
		return
	}
	a.mailbox <- v
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

func (a *pingActor) PreStart() {

}

func NewPingActor(name string) actor.ActorRef {
	return &pingActor{
		BaseActor: actor.BaseActor{Name: name, Mailbox: make(chan any)},
		mailbox:   make(chan string),
	}
}
