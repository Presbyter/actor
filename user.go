package actor

type User struct {
	name   string
	actors map[string]ActorRef
	System *System
}

func (u *User) RegisterActor(actors ...ActorRef) {
	for _, actor := range actors {
		name := actor.GetName()
		if _, ok := u.actors[name]; ok {
			continue
		}
		u.actors[name] = actor
		actor.SetUser(u)
		actor.PreStart()
		go actor.Start()
	}
}

func (u *User) GetActor(name string) ActorRef {
	v, ok := u.actors[name]
	if !ok {
		return nil
	}
	return v
}

func (u *User) setSystem(s *System) {
	u.System = s
}

func NewUser(name string) *User {
	return &User{name: name, actors: make(map[string]ActorRef)}
}
