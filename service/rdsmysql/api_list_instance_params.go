// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListInstanceParamsCommon = "ListInstanceParams"

// ListInstanceParamsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListInstanceParamsCommon operation. The "output" return
// value will be populated with the ListInstanceParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListInstanceParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListInstanceParamsCommon Send returns without error.
//
// See ListInstanceParamsCommon for more information on using the ListInstanceParamsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListInstanceParamsCommonRequest method.
//    req, resp := client.ListInstanceParamsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListInstanceParamsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListInstanceParamsCommon,
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

// ListInstanceParamsCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListInstanceParamsCommon for usage and error information.
func (c *RDSMYSQL) ListInstanceParamsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListInstanceParamsCommonRequest(input)
	return out, req.Send()
}

// ListInstanceParamsCommonWithContext is the same as ListInstanceParamsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListInstanceParamsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListInstanceParamsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListInstanceParamsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListInstanceParams = "ListInstanceParams"

// ListInstanceParamsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListInstanceParams operation. The "output" return
// value will be populated with the ListInstanceParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListInstanceParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListInstanceParamsCommon Send returns without error.
//
// See ListInstanceParams for more information on using the ListInstanceParams
// API call, and error handling.
//
//    // Example sending a request using the ListInstanceParamsRequest method.
//    req, resp := client.ListInstanceParamsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListInstanceParamsRequest(input *ListInstanceParamsInput) (req *request.Request, output *ListInstanceParamsOutput) {
	op := &request.Operation{
		Name:       opListInstanceParams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListInstanceParamsInput{}
	}

	output = &ListInstanceParamsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListInstanceParams API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListInstanceParams for usage and error information.
func (c *RDSMYSQL) ListInstanceParams(input *ListInstanceParamsInput) (*ListInstanceParamsOutput, error) {
	req, out := c.ListInstanceParamsRequest(input)
	return out, req.Send()
}

// ListInstanceParamsWithContext is the same as ListInstanceParams with the addition of
// the ability to pass a context and additional request options.
//
// See ListInstanceParams for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListInstanceParamsWithContext(ctx volcengine.Context, input *ListInstanceParamsInput, opts ...request.Option) (*ListInstanceParamsOutput, error) {
	req, out := c.ListInstanceParamsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListInstanceParamsOutput struct {
	_ struct{} `type:"structure"`

	DefaultValue *string `type:"string"`

	Description *string `type:"string"`

	ExpectValue *string `type:"string"`

	Name *string `type:"string"`

	Restart *bool `type:"boolean"`

	RunningValue *string `type:"string"`

	ValueRange *string `type:"string"`
}

// String returns the string representation
func (s DataForListInstanceParamsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListInstanceParamsOutput) GoString() string {
	return s.String()
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *DataForListInstanceParamsOutput) SetDefaultValue(v string) *DataForListInstanceParamsOutput {
	s.DefaultValue = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListInstanceParamsOutput) SetDescription(v string) *DataForListInstanceParamsOutput {
	s.Description = &v
	return s
}

// SetExpectValue sets the ExpectValue field's value.
func (s *DataForListInstanceParamsOutput) SetExpectValue(v string) *DataForListInstanceParamsOutput {
	s.ExpectValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListInstanceParamsOutput) SetName(v string) *DataForListInstanceParamsOutput {
	s.Name = &v
	return s
}

// SetRestart sets the Restart field's value.
func (s *DataForListInstanceParamsOutput) SetRestart(v bool) *DataForListInstanceParamsOutput {
	s.Restart = &v
	return s
}

// SetRunningValue sets the RunningValue field's value.
func (s *DataForListInstanceParamsOutput) SetRunningValue(v string) *DataForListInstanceParamsOutput {
	s.RunningValue = &v
	return s
}

// SetValueRange sets the ValueRange field's value.
func (s *DataForListInstanceParamsOutput) SetValueRange(v string) *DataForListInstanceParamsOutput {
	s.ValueRange = &v
	return s
}

type ListInstanceParamsInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListInstanceParamsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListInstanceParamsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInstanceParamsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListInstanceParamsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListInstanceParamsInput) SetInstanceId(v string) *ListInstanceParamsInput {
	s.InstanceId = &v
	return s
}

type ListInstanceParamsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Datas []*DataForListInstanceParamsOutput `type:"list"`
}

// String returns the string representation
func (s ListInstanceParamsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListInstanceParamsOutput) GoString() string {
	return s.String()
}

// SetDatas sets the Datas field's value.
func (s *ListInstanceParamsOutput) SetDatas(v []*DataForListInstanceParamsOutput) *ListInstanceParamsOutput {
	s.Datas = v
	return s
}
