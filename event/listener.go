package event

type Listener interface {
	Handle(event Event)
}
