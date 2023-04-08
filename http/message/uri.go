package message

type Uri interface {
	GetScheme() string
	GetAuthority() string
	GetUserInfo() string
	GetHost() string
	GetPort() int
	GetPath() string
	GetQuery() string
	GetFragment() string
	WithScheme(string) Uri
	WithUserInfo(string, string) Uri
	WithHost(string) Uri
	WithPort(int) Uri
	WithPath(string) Uri
	WithQuery(string) Uri
	WithFragment(string) Uri
	String() string
}
