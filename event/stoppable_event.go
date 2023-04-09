package event

type StoppableEvent interface {
	IsPropagationStopped() bool
}
