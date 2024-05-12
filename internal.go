package actor

import "log"

type eventbusActor struct {
	BaseActor
}

func (a *eventbusActor) Start() {
	for msg := range a.Mailbox {
		// transmit to subscriber
		log.Println("eventActor transmit: ", msg)
		// a.User.System.
	}
}

func NewEventbusActor(name string) ActorRef {
	return &eventbusActor{
		BaseActor: BaseActor{Name: name, Mailbox: make(chan any)},
	}
}
