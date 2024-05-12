package main

import (
	"time"

	"github.com/Presbyter/actor"
	"github.com/Presbyter/actor/example/actors"
)

func main() {
	actorSystem := actor.NewSystem("PingPong",
		// actor.WithUser("user1"),
		actor.WithActor("user1", actors.NewPingActor("ping")),
		actor.WithActor("user1", actors.NewPongActor("pong")),
	)

	arr := []string{"1", "2", "3"}
	for _, v := range arr {
		a := actorSystem.GetActor("user1/ping")
		if a == nil {
			continue
		}
		a.Send(v)
	}

	time.Sleep(time.Second)
}
