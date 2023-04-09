package event

type NullDispatcher struct {
	listeners map[string][]Listener
}

var _ Dispatcher = (*NullDispatcher)(nil)

func NewNullDispatcher() *NullDispatcher {
	return &NullDispatcher{
		listeners: make(map[string][]Listener, 0),
	}
}

func (n NullDispatcher) Listen(name string, listener Listener) {
	n.listeners[name] = append(n.listeners[name], listener)
}

func (n NullDispatcher) Dispatch(event Event) {
	for _, listener := range n.GetListeners(event.Name()) {
		if event.(StoppableEvent) != nil && event.(StoppableEvent).IsPropagationStopped() {
			return
		}

		listener.Handle(event)
	}
}

func (n NullDispatcher) GetListeners(name string) []Listener {
	if listeners, ok := n.listeners[name]; ok {
		return listeners
	}

	return nil
}

func (n NullDispatcher) Flush(name string) {
	if _, ok := n.listeners[name]; ok {
		delete(n.listeners, name)
	}
}

func (n NullDispatcher) FlushAll() {
	n.listeners = make(map[string][]Listener, 0)
}
