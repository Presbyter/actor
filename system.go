package actor

import (
	"strings"
)

type System struct {
	name     string
	users    map[string]*User
	eventbus chan any
}

func (s *System) RegisterUser(users ...*User) {
	for _, user := range users {
		name := user.name
		s.users[name] = user
	}
}

func (s *System) GetUser(name string) *User {
	v, ok := s.users[name]
	if !ok {
		return nil
	}
	return v
}

func (s *System) GetActor(path string) ActorRef {
	arr := strings.SplitN(path, "/", 2)
	userName, actorName := arr[0], arr[1]
	user := s.GetUser(userName)
	if user == nil {
		return nil
	}
	actor := user.GetActor(actorName)
	return actor
}

func NewSystem(name string, opts ...ConfigOption) *System {
	s := &System{name: name,
		users:    make(map[string]*User),
		eventbus: make(chan any),
	}
	for _, opt := range opts {
		opt(s)
	}
	systemAgent[name] = s
	return s
}

// ------------------------------------------------------------------

type ConfigOption func(*System)

func WithUser(name string) ConfigOption {
	return func(s *System) {
		s.RegisterUser(NewUser(name))
	}
}

func WithActor(userName string, actor ActorRef) ConfigOption {
	return func(s *System) {
		// check user is exist
		user := s.GetUser(userName)
		if user == nil {
			s.RegisterUser(NewUser(userName))
			user = s.GetUser(userName)
		}

		user.RegisterActor(actor)
	}
}

// ------------------------------------------------------------------

var systemAgent map[string]*System

func init() {
	systemAgent = make(map[string]*System)
}

func GetSystem(name string) (*System, bool) {
	v, ok := systemAgent[name]
	return v, ok
}

func GetActor(path string) (ActorRef, bool) {
	arr := strings.SplitN(path, "/", 3)
	systemName, userName, actorName := arr[0], arr[1], arr[2]
	s, ok := GetSystem(systemName)
	if !ok {
		return nil, ok
	}
	u := s.GetUser(userName)
	if u == nil {
		return nil, false
	}
	a := u.GetActor(actorName)
	if a == nil {
		return nil, false
	}
	return a, true
}
