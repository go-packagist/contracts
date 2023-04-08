package message

type Request interface {
	Message

	GetRequestTarget() string
	WithRequestTarget(requestTarget string) Request
	GetMethod() string
	WithMethod(method string) Request
	GetUri() Uri
	WithUri(uri Uri, preserveHost bool) Request
}

type ServerRequest interface {
	Request

	GetServerParams() map[string]interface{}
	GetCookieParams() map[string]string
	WithCookieParams(cookies map[string]string) ServerRequest
	GetQueryParams() map[string][]string
	WithQueryParams(params map[string][]string) ServerRequest
	GetUploadedFiles() map[string]UploadedFile
	WithUploadedFiles(files map[string]UploadedFile) ServerRequest
	GetParsedBody() interface{}
}
