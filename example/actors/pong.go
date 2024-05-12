package actors

import (
	"fmt"

	"github.com/Presbyter/actor"
)

type pongActor struct {
	name    string
	mailbox chan any
}

func (a *pongActor) Name() string {
	return a.name
}

func (a *pongActor) Send(msg any) {
	a.mailbox <- msg
}

func (a *pongActor) Start() {
	for msg := range a.mailbox {
		fmt.Println("PONG. Get msg:", msg)
	}
}

func NewPongActor(name string) actor.ActorRef {
	return &pongActor{
		name:    name,
		mailbox: make(chan any),
	}
}
