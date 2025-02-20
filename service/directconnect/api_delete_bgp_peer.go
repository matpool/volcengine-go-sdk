// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteBgpPeerCommon = "DeleteBgpPeer"

// DeleteBgpPeerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteBgpPeerCommon operation. The "output" return
// value will be populated with the DeleteBgpPeerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteBgpPeerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteBgpPeerCommon Send returns without error.
//
// See DeleteBgpPeerCommon for more information on using the DeleteBgpPeerCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteBgpPeerCommonRequest method.
//	req, resp := client.DeleteBgpPeerCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DeleteBgpPeerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteBgpPeerCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteBgpPeerCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DeleteBgpPeerCommon for usage and error information.
func (c *DIRECTCONNECT) DeleteBgpPeerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteBgpPeerCommonRequest(input)
	return out, req.Send()
}

// DeleteBgpPeerCommonWithContext is the same as DeleteBgpPeerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBgpPeerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DeleteBgpPeerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteBgpPeerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteBgpPeer = "DeleteBgpPeer"

// DeleteBgpPeerRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteBgpPeer operation. The "output" return
// value will be populated with the DeleteBgpPeerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteBgpPeerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteBgpPeerCommon Send returns without error.
//
// See DeleteBgpPeer for more information on using the DeleteBgpPeer
// API call, and error handling.
//
//	// Example sending a request using the DeleteBgpPeerRequest method.
//	req, resp := client.DeleteBgpPeerRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DeleteBgpPeerRequest(input *DeleteBgpPeerInput) (req *request.Request, output *DeleteBgpPeerOutput) {
	op := &request.Operation{
		Name:       opDeleteBgpPeer,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBgpPeerInput{}
	}

	output = &DeleteBgpPeerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteBgpPeer API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DeleteBgpPeer for usage and error information.
func (c *DIRECTCONNECT) DeleteBgpPeer(input *DeleteBgpPeerInput) (*DeleteBgpPeerOutput, error) {
	req, out := c.DeleteBgpPeerRequest(input)
	return out, req.Send()
}

// DeleteBgpPeerWithContext is the same as DeleteBgpPeer with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBgpPeer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DeleteBgpPeerWithContext(ctx volcengine.Context, input *DeleteBgpPeerInput, opts ...request.Option) (*DeleteBgpPeerOutput, error) {
	req, out := c.DeleteBgpPeerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteBgpPeerInput struct {
	_ struct{} `type:"structure"`

	// BgpPeerId is a required field
	BgpPeerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBgpPeerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBgpPeerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBgpPeerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBgpPeerInput"}
	if s.BgpPeerId == nil {
		invalidParams.Add(request.NewErrParamRequired("BgpPeerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBgpPeerId sets the BgpPeerId field's value.
func (s *DeleteBgpPeerInput) SetBgpPeerId(v string) *DeleteBgpPeerInput {
	s.BgpPeerId = &v
	return s
}

type DeleteBgpPeerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteBgpPeerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBgpPeerOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteBgpPeerOutput) SetRequestId(v string) *DeleteBgpPeerOutput {
	s.RequestId = &v
	return s
}
