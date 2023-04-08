package message

type Message interface {
	GetProtocolVersion() string
	withProtocolVersion(string) Message
	GetHeaders() map[string][]string
	HasHeader(string) bool
	GetHeader(string) []string
	GetHeaderLine(string) string
	WithHeader(string, string) Message
	WithAddedHeader(string, string) Message
	WithoutHeader(string) Message
	GetBody() Stream
	WithBody(Stream) Message
}
