package endpointcreds

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"encoding/json"
	"time"

	"github.com/matpool/volcengine-go-sdk/private/protocol/json/jsonutil"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/client"
	"github.com/matpool/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineerr"
)

// ProviderName is the name of the credentials provider.
const ProviderName = `CredentialsEndpointProvider`

// Provider satisfies the credentials.Provider interface, and is a client to
// retrieve credentials from an arbitrary endpoint.
type Provider struct {
	staticCreds bool
	credentials.Expiry

	// Requires a Volcengine Client to make HTTP requests to the endpoint with.
	// the Endpoint the request will be made to is provided by the volcengine.Config's
	// Endpoint value.
	Client *client.Client

	// ExpiryWindow will allow the credentials to trigger refreshing prior to
	// the credentials actually expiring. This is beneficial so race conditions
	// with expiring credentials do not cause request to fail unexpectedly
	// due to ExpiredTokenException exceptions.
	//
	// So a ExpiryWindow of 10s would cause calls to IsExpired() to return true
	// 10 seconds before the credentials are actually expired.
	//
	// If ExpiryWindow is 0 or less it will be ignored.
	ExpiryWindow time.Duration

	// Optional authorization token value if set will be used as the value of
	// the Authorization header of the endpoint credential request.
	AuthorizationToken string
}

// NewProviderClient returns a credentials Provider for retrieving Volcengine credentials
// from arbitrary endpoint.
func NewProviderClient(cfg volcengine.Config, handlers request.Handlers, endpoint string, options ...func(*Provider)) credentials.Provider {
	p := &Provider{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName: "CredentialsEndpoint",
				Endpoint:    endpoint,
			},
			handlers,
		),
	}

	p.Client.Handlers.Unmarshal.PushBack(unmarshalHandler)
	p.Client.Handlers.UnmarshalError.PushBack(unmarshalError)
	p.Client.Handlers.Validate.Clear()
	p.Client.Handlers.Validate.PushBack(validateEndpointHandler)

	for _, option := range options {
		option(p)
	}

	return p
}

// NewCredentialsClient returns a pointer to a new Credentials object
// wrapping the endpoint credentials Provider.
func NewCredentialsClient(cfg volcengine.Config, handlers request.Handlers, endpoint string, options ...func(*Provider)) *credentials.Credentials {
	return credentials.NewCredentials(NewProviderClient(cfg, handlers, endpoint, options...))
}

// IsExpired returns true if the credentials retrieved are expired, or not yet
// retrieved.
func (p *Provider) IsExpired() bool {
	if p.staticCreds {
		return false
	}
	return p.Expiry.IsExpired()
}

// Retrieve will attempt to request the credentials from the endpoint the Provider
// was configured for. And error will be returned if the retrieval fails.
func (p *Provider) Retrieve() (credentials.Value, error) {
	resp, err := p.getCredentials()
	if err != nil {
		return credentials.Value{ProviderName: ProviderName},
			volcengineerr.New("CredentialsEndpointError", "failed to load credentials", err)
	}

	if resp.Expiration != nil {
		p.SetExpiration(*resp.Expiration, p.ExpiryWindow)
	} else {
		p.staticCreds = true
	}

	return credentials.Value{
		AccessKeyID:     resp.AccessKeyID,
		SecretAccessKey: resp.SecretAccessKey,
		SessionToken:    resp.Token,
		ProviderName:    ProviderName,
	}, nil
}

type getCredentialsOutput struct {
	Expiration      *time.Time
	AccessKeyID     string
	SecretAccessKey string
	Token           string
}

type errorOutput struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (p *Provider) getCredentials() (*getCredentialsOutput, error) {
	op := &request.Operation{
		Name:       "GetCredentials",
		HTTPMethod: "GET",
	}

	out := &getCredentialsOutput{}
	req := p.Client.NewRequest(op, nil, out)
	req.HTTPRequest.Header.Set("Accept", "application/json")
	if authToken := p.AuthorizationToken; len(authToken) != 0 {
		req.HTTPRequest.Header.Set("Authorization", authToken)
	}

	return out, req.Send()
}

func validateEndpointHandler(r *request.Request) {
	if len(r.ClientInfo.Endpoint) == 0 {
		r.Error = volcengine.ErrMissingEndpoint
	}
}

func unmarshalHandler(r *request.Request) {
	defer r.HTTPResponse.Body.Close()

	out := r.Data.(*getCredentialsOutput)
	if err := json.NewDecoder(r.HTTPResponse.Body).Decode(&out); err != nil {
		r.Error = volcengineerr.New(request.ErrCodeSerialization,
			"failed to decode endpoint credentials",
			err,
		)
	}
}

func unmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()

	var errOut errorOutput
	err := jsonutil.UnmarshalJSONError(&errOut, r.HTTPResponse.Body)
	if err != nil {
		r.Error = volcengineerr.NewRequestFailure(
			volcengineerr.New(request.ErrCodeSerialization,
				"failed to decode error message", err),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}

	// Response volcenginebody format is not consistent between metadata endpoints.
	// Grab the error message as a string and include that as the source error
	r.Error = volcengineerr.New(errOut.Code, errOut.Message, nil)
}
