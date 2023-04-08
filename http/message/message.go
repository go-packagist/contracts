package message

type Message interface {
	GetProtocolVersion() string
	withProtocolVersion(version string) Message
	GetHeaders() map[string][]string
	HasHeader(name string) bool
	GetHeader(name string) []string
	GetHeaderLine(name string) string
	WithHeader(name string, value ...string) Message
	WithAddedHeader(name string, value ...string) Message
	WithoutHeader(name string) Message
	GetBody() Stream
	WithBody(body Stream) Message
}
