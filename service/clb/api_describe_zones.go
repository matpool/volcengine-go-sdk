// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeZonesCommon = "DescribeZones"

// DescribeZonesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeZonesCommon operation. The "output" return
// value will be populated with the DescribeZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeZonesCommon Send returns without error.
//
// See DescribeZonesCommon for more information on using the DescribeZonesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeZonesCommonRequest method.
//    req, resp := client.DescribeZonesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeZonesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeZonesCommon,
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

// DescribeZonesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeZonesCommon for usage and error information.
func (c *CLB) DescribeZonesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeZonesCommonRequest(input)
	return out, req.Send()
}

// DescribeZonesCommonWithContext is the same as DescribeZonesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeZonesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeZonesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeZonesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeZones = "DescribeZones"

// DescribeZonesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeZones operation. The "output" return
// value will be populated with the DescribeZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeZonesCommon Send returns without error.
//
// See DescribeZones for more information on using the DescribeZones
// API call, and error handling.
//
//    // Example sending a request using the DescribeZonesRequest method.
//    req, resp := client.DescribeZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeZonesRequest(input *DescribeZonesInput) (req *request.Request, output *DescribeZonesOutput) {
	op := &request.Operation{
		Name:       opDescribeZones,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeZonesInput{}
	}

	output = &DescribeZonesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeZones API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeZones for usage and error information.
func (c *CLB) DescribeZones(input *DescribeZonesInput) (*DescribeZonesOutput, error) {
	req, out := c.DescribeZonesRequest(input)
	return out, req.Send()
}

// DescribeZonesWithContext is the same as DescribeZones with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeZonesWithContext(ctx volcengine.Context, input *DescribeZonesInput, opts ...request.Option) (*DescribeZonesOutput, error) {
	req, out := c.DescribeZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeZonesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeZonesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeZonesInput) GoString() string {
	return s.String()
}

type DescribeZonesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	MasterZones []*MasterZoneForDescribeZonesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeZonesOutput) GoString() string {
	return s.String()
}

// SetMasterZones sets the MasterZones field's value.
func (s *DescribeZonesOutput) SetMasterZones(v []*MasterZoneForDescribeZonesOutput) *DescribeZonesOutput {
	s.MasterZones = v
	return s
}

type MasterZoneForDescribeZonesOutput struct {
	_ struct{} `type:"structure"`

	SlaveZones []*string `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s MasterZoneForDescribeZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MasterZoneForDescribeZonesOutput) GoString() string {
	return s.String()
}

// SetSlaveZones sets the SlaveZones field's value.
func (s *MasterZoneForDescribeZonesOutput) SetSlaveZones(v []*string) *MasterZoneForDescribeZonesOutput {
	s.SlaveZones = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *MasterZoneForDescribeZonesOutput) SetZoneId(v string) *MasterZoneForDescribeZonesOutput {
	s.ZoneId = &v
	return s
}
