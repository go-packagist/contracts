package message

type Stream interface {
	String() string
	Close() error

	// todo: add more methods
}
