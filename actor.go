package actor

type ActorRef interface {
	Send(msg any)
	Name() string
	Start()
}
