package universal

import (
	"fmt"
	"reflect"

	"github.com/matpool/volcengine-go-sdk/volcengine/client"
	"github.com/matpool/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/matpool/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
	"github.com/matpool/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcenginequery"
)

func New(session *session.Session) *Universal {
	return &Universal{
		Session: session,
	}
}

func (u *Universal) newClient(info RequestUniversal) *client.Client {
	config := u.Session.ClientConfig(info.ServiceName)
	c := client.New(
		*config.Config,
		metadata.ClientInfo{
			SigningName:   config.SigningName,
			SigningRegion: config.SigningRegion,
			Endpoint:      config.Endpoint,
			APIVersion:    info.Version,
			ServiceName:   info.ServiceName,
			ServiceID:     info.ServiceName,
		},
		config.Handlers,
	)
	c.Handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	c.Handlers.Sign.PushBackNamed(volc.SignRequestHandler)
	c.Handlers.Build.PushBackNamed(volcenginequery.BuildHandler)
	c.Handlers.Unmarshal.PushBackNamed(volcenginequery.UnmarshalHandler)
	c.Handlers.UnmarshalMeta.PushBackNamed(volcenginequery.UnmarshalMetaHandler)
	c.Handlers.UnmarshalError.PushBackNamed(volcenginequery.UnmarshalErrorHandler)

	return c
}

func (u *Universal) getMethod(m HttpMethod) string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case HEAD:
		return "HEAD"
	default:
		return "GET"
	}
}

func getContentType(m ContentType) string {
	switch m {
	case ApplicationJSON:
		return "application/json"
	case FormUrlencoded:
		return "x-www-form-urlencoded"
	default:
		return ""
	}
}

func (u *Universal) DoCall(info RequestUniversal, input *map[string]interface{}) (output *map[string]interface{}, err error) {
	c := u.newClient(info)
	op := &request.Operation{
		HTTPMethod: u.getMethod(info.HttpMethod),
		HTTPPath:   "/",
		Name:       info.Action,
	}
	if input == nil {
		input = &map[string]interface{}{}
	}
	output = &map[string]interface{}{}
	req := c.NewRequest(op, input, output)

	if getContentType(info.ContentType) != "" {
		req.HTTPRequest.Header.Set("Content-Type", getContentType(info.ContentType))
	}
	err = req.Send()
	return output, err
}

func (u *Universal) DoCallWithType(info RequestUniversal, input interface{}, output interface{}) (err error) {
	c := u.newClient(info)
	op := &request.Operation{
		HTTPMethod: u.getMethod(info.HttpMethod),
		HTTPPath:   "/",
		Name:       info.Action,
	}
	if input == nil {
		input = &map[string]interface{}{}
	} else if reflect.TypeOf(input).Kind() != reflect.Ptr {
		return fmt.Errorf("input is not pointor ")
	}
	if output == nil {
		output = &map[string]interface{}{}
	} else if reflect.TypeOf(output).Kind() != reflect.Ptr {
		return fmt.Errorf("output is not pointor ")
	}
	req := c.NewRequest(op, input, output)
	if getContentType(info.ContentType) != "" {
		req.HTTPRequest.Header.Set("Content-Type", getContentType(info.ContentType))
	}
	return req.Send()
}
