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
	WithScheme(scheme string) Uri
	WithUserInfo(user string, password ...string) Uri
	WithHost(host string) Uri
	WithPort(port int) Uri
	WithPath(path string) Uri
	WithQuery(query string) Uri
	WithFragment(fragment string) Uri
	String() string
}
