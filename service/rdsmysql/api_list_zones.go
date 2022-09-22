// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListZonesCommon = "ListZones"

// ListZonesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListZonesCommon operation. The "output" return
// value will be populated with the ListZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListZonesCommon Send returns without error.
//
// See ListZonesCommon for more information on using the ListZonesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListZonesCommonRequest method.
//    req, resp := client.ListZonesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListZonesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListZonesCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListZonesCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListZonesCommon for usage and error information.
func (c *RDSMYSQL) ListZonesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListZonesCommonRequest(input)
	return out, req.Send()
}

// ListZonesCommonWithContext is the same as ListZonesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListZonesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListZonesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListZonesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListZones = "ListZones"

// ListZonesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListZones operation. The "output" return
// value will be populated with the ListZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListZonesCommon Send returns without error.
//
// See ListZones for more information on using the ListZones
// API call, and error handling.
//
//    // Example sending a request using the ListZonesRequest method.
//    req, resp := client.ListZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListZonesRequest(input *ListZonesInput) (req *request.Request, output *ListZonesOutput) {
	op := &request.Operation{
		Name:       opListZones,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListZonesInput{}
	}

	output = &ListZonesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListZones API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListZones for usage and error information.
func (c *RDSMYSQL) ListZones(input *ListZonesInput) (*ListZonesOutput, error) {
	req, out := c.ListZonesRequest(input)
	return out, req.Send()
}

// ListZonesWithContext is the same as ListZones with the addition of
// the ability to pass a context and additional request options.
//
// See ListZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListZonesWithContext(ctx volcengine.Context, input *ListZonesInput, opts ...request.Option) (*ListZonesOutput, error) {
	req, out := c.ListZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListZonesInput struct {
	_ struct{} `type:"structure"`

	// Region is a required field
	Region *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListZonesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListZonesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListZonesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListZonesInput"}
	if s.Region == nil {
		invalidParams.Add(request.NewErrParamRequired("Region"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRegion sets the Region field's value.
func (s *ListZonesInput) SetRegion(v string) *ListZonesInput {
	s.Region = &v
	return s
}

type ListZonesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ListZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListZonesOutput) GoString() string {
	return s.String()
}
