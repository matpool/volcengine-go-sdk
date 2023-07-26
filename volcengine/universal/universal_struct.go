package universal

import "github.com/matpool/volcengine-go-sdk/volcengine/session"

type Universal struct {
	Session *session.Session
}

type RequestUniversal struct {
	ServiceName string
	Action      string
	Version     string
	HttpMethod  HttpMethod
	ContentType ContentType
}
