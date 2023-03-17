// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCenSummaryRouteEntryCommon = "DeleteCenSummaryRouteEntry"

// DeleteCenSummaryRouteEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCenSummaryRouteEntryCommon operation. The "output" return
// value will be populated with the DeleteCenSummaryRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCenSummaryRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCenSummaryRouteEntryCommon Send returns without error.
//
// See DeleteCenSummaryRouteEntryCommon for more information on using the DeleteCenSummaryRouteEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCenSummaryRouteEntryCommonRequest method.
//    req, resp := client.DeleteCenSummaryRouteEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DeleteCenSummaryRouteEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCenSummaryRouteEntryCommon,
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

// DeleteCenSummaryRouteEntryCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DeleteCenSummaryRouteEntryCommon for usage and error information.
func (c *CEN) DeleteCenSummaryRouteEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCenSummaryRouteEntryCommonRequest(input)
	return out, req.Send()
}

// DeleteCenSummaryRouteEntryCommonWithContext is the same as DeleteCenSummaryRouteEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCenSummaryRouteEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DeleteCenSummaryRouteEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCenSummaryRouteEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCenSummaryRouteEntry = "DeleteCenSummaryRouteEntry"

// DeleteCenSummaryRouteEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCenSummaryRouteEntry operation. The "output" return
// value will be populated with the DeleteCenSummaryRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCenSummaryRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCenSummaryRouteEntryCommon Send returns without error.
//
// See DeleteCenSummaryRouteEntry for more information on using the DeleteCenSummaryRouteEntry
// API call, and error handling.
//
//    // Example sending a request using the DeleteCenSummaryRouteEntryRequest method.
//    req, resp := client.DeleteCenSummaryRouteEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DeleteCenSummaryRouteEntryRequest(input *DeleteCenSummaryRouteEntryInput) (req *request.Request, output *DeleteCenSummaryRouteEntryOutput) {
	op := &request.Operation{
		Name:       opDeleteCenSummaryRouteEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCenSummaryRouteEntryInput{}
	}

	output = &DeleteCenSummaryRouteEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCenSummaryRouteEntry API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DeleteCenSummaryRouteEntry for usage and error information.
func (c *CEN) DeleteCenSummaryRouteEntry(input *DeleteCenSummaryRouteEntryInput) (*DeleteCenSummaryRouteEntryOutput, error) {
	req, out := c.DeleteCenSummaryRouteEntryRequest(input)
	return out, req.Send()
}

// DeleteCenSummaryRouteEntryWithContext is the same as DeleteCenSummaryRouteEntry with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCenSummaryRouteEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DeleteCenSummaryRouteEntryWithContext(ctx volcengine.Context, input *DeleteCenSummaryRouteEntryInput, opts ...request.Option) (*DeleteCenSummaryRouteEntryOutput, error) {
	req, out := c.DeleteCenSummaryRouteEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCenSummaryRouteEntryInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`

	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCenSummaryRouteEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCenSummaryRouteEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCenSummaryRouteEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCenSummaryRouteEntryInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}
	if s.DestinationCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationCidrBlock"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *DeleteCenSummaryRouteEntryInput) SetCenId(v string) *DeleteCenSummaryRouteEntryInput {
	s.CenId = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DeleteCenSummaryRouteEntryInput) SetDestinationCidrBlock(v string) *DeleteCenSummaryRouteEntryInput {
	s.DestinationCidrBlock = &v
	return s
}

type DeleteCenSummaryRouteEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteCenSummaryRouteEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCenSummaryRouteEntryOutput) GoString() string {
	return s.String()
}
