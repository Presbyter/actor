package actors

import (
	"fmt"

	"github.com/Presbyter/actor"
)

type pongActor struct {
	actor.BaseActor
}

func (a *pongActor) Start() {
	for msg := range a.Mailbox {
		fmt.Println("PONG. Get msg:", msg)
	}
}

func (a *pongActor) PreStart() {

}

func NewPongActor(name string) actor.ActorRef {
	return &pongActor{
		BaseActor: actor.BaseActor{
			Name:    name,
			Mailbox: make(chan any),
		},
	}
}
