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

// ------------------------------------------------------------------

type Event[T any] struct {
	From    string
	To      string
	Forward string
	Data    T
}

func (e *Event[T]) GetData() T {
	return e.Data
}

func NewEvent[T any](data T, opts ...EventOption[T]) Event[T] {
	e := Event[T]{Data: data}
	for _, opt := range opts {
		opt(&e)
	}
	return e
}

type EventOption[T any] func(*Event[T])

func WithFrom[T any](from string) EventOption[T] {
	return func(e *Event[T]) {
		e.From = from
	}
}

func WithTo[T any](to string) EventOption[T] {
	return func(e *Event[T]) {
		e.To = to
	}
}

func WithForward[T any](forward string) EventOption[T] {
	return func(e *Event[T]) {
		e.Forward = forward
	}
}
