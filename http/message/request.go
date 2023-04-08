package message

type Request interface {
	Message

	GetRequestTarget() string
	WithRequestTarget(string) Request
	GetMethod() string
	WithMethod(string) Request
	GetUri() Uri
}
