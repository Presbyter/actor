package actor

type ActorRef interface {
	Send(msg any)
	GetName() string
	Start()
	PreStart()

	SetUser(user *User)
}

type BaseActor struct {
	Name    string
	User    *User
	Mailbox chan any
}

func (a *BaseActor) Send(msg any) {
	if a.Mailbox == nil {
		a.Mailbox = make(chan any)
	}
	a.Mailbox <- msg
}

func (a *BaseActor) GetName() string {
	return a.Name
}

func (a *BaseActor) Start() {
	panic("please implement me")
}

func (a *BaseActor) PreStart() {

}

func (a *BaseActor) SetUser(user *User) {
	a.User = user
}
