// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNodeSpecsCommon = "DescribeNodeSpecs"

// DescribeNodeSpecsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNodeSpecsCommon operation. The "output" return
// value will be populated with the DescribeNodeSpecsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNodeSpecsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNodeSpecsCommon Send returns without error.
//
// See DescribeNodeSpecsCommon for more information on using the DescribeNodeSpecsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeNodeSpecsCommonRequest method.
//    req, resp := client.DescribeNodeSpecsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeNodeSpecsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNodeSpecsCommon,
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

// DescribeNodeSpecsCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeNodeSpecsCommon for usage and error information.
func (c *MONGODB) DescribeNodeSpecsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNodeSpecsCommonRequest(input)
	return out, req.Send()
}

// DescribeNodeSpecsCommonWithContext is the same as DescribeNodeSpecsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNodeSpecsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeNodeSpecsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNodeSpecsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNodeSpecs = "DescribeNodeSpecs"

// DescribeNodeSpecsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNodeSpecs operation. The "output" return
// value will be populated with the DescribeNodeSpecsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNodeSpecsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNodeSpecsCommon Send returns without error.
//
// See DescribeNodeSpecs for more information on using the DescribeNodeSpecs
// API call, and error handling.
//
//    // Example sending a request using the DescribeNodeSpecsRequest method.
//    req, resp := client.DescribeNodeSpecsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeNodeSpecsRequest(input *DescribeNodeSpecsInput) (req *request.Request, output *DescribeNodeSpecsOutput) {
	op := &request.Operation{
		Name:       opDescribeNodeSpecs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNodeSpecsInput{}
	}

	output = &DescribeNodeSpecsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeNodeSpecs API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeNodeSpecs for usage and error information.
func (c *MONGODB) DescribeNodeSpecs(input *DescribeNodeSpecsInput) (*DescribeNodeSpecsOutput, error) {
	req, out := c.DescribeNodeSpecsRequest(input)
	return out, req.Send()
}

// DescribeNodeSpecsWithContext is the same as DescribeNodeSpecs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNodeSpecs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeNodeSpecsWithContext(ctx volcengine.Context, input *DescribeNodeSpecsInput, opts ...request.Option) (*DescribeNodeSpecsOutput, error) {
	req, out := c.DescribeNodeSpecsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeNodeSpecsInput struct {
	_ struct{} `type:"structure"`

	RegionId *string `type:"string"`
}

// String returns the string representation
func (s DescribeNodeSpecsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNodeSpecsInput) GoString() string {
	return s.String()
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeNodeSpecsInput) SetRegionId(v string) *DescribeNodeSpecsInput {
	s.RegionId = &v
	return s
}

type DescribeNodeSpecsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	MongosNodeSpecs []*MongosNodeSpecForDescribeNodeSpecsOutput `type:"list"`

	NodeSpecs []*NodeSpecForDescribeNodeSpecsOutput `type:"list"`

	ShardNodeSpecs []*ShardNodeSpecForDescribeNodeSpecsOutput `type:"list"`
}

// String returns the string representation
func (s DescribeNodeSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNodeSpecsOutput) GoString() string {
	return s.String()
}

// SetMongosNodeSpecs sets the MongosNodeSpecs field's value.
func (s *DescribeNodeSpecsOutput) SetMongosNodeSpecs(v []*MongosNodeSpecForDescribeNodeSpecsOutput) *DescribeNodeSpecsOutput {
	s.MongosNodeSpecs = v
	return s
}

// SetNodeSpecs sets the NodeSpecs field's value.
func (s *DescribeNodeSpecsOutput) SetNodeSpecs(v []*NodeSpecForDescribeNodeSpecsOutput) *DescribeNodeSpecsOutput {
	s.NodeSpecs = v
	return s
}

// SetShardNodeSpecs sets the ShardNodeSpecs field's value.
func (s *DescribeNodeSpecsOutput) SetShardNodeSpecs(v []*ShardNodeSpecForDescribeNodeSpecsOutput) *DescribeNodeSpecsOutput {
	s.ShardNodeSpecs = v
	return s
}

type MongosNodeSpecForDescribeNodeSpecsOutput struct {
	_ struct{} `type:"structure"`

	CpuNum *float64 `type:"double"`

	MaxConn *int64 `type:"int64"`

	MemInGb *float64 `type:"double"`

	SpecName *string `type:"string"`
}

// String returns the string representation
func (s MongosNodeSpecForDescribeNodeSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MongosNodeSpecForDescribeNodeSpecsOutput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *MongosNodeSpecForDescribeNodeSpecsOutput) SetCpuNum(v float64) *MongosNodeSpecForDescribeNodeSpecsOutput {
	s.CpuNum = &v
	return s
}

// SetMaxConn sets the MaxConn field's value.
func (s *MongosNodeSpecForDescribeNodeSpecsOutput) SetMaxConn(v int64) *MongosNodeSpecForDescribeNodeSpecsOutput {
	s.MaxConn = &v
	return s
}

// SetMemInGb sets the MemInGb field's value.
func (s *MongosNodeSpecForDescribeNodeSpecsOutput) SetMemInGb(v float64) *MongosNodeSpecForDescribeNodeSpecsOutput {
	s.MemInGb = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *MongosNodeSpecForDescribeNodeSpecsOutput) SetSpecName(v string) *MongosNodeSpecForDescribeNodeSpecsOutput {
	s.SpecName = &v
	return s
}

type NodeSpecForDescribeNodeSpecsOutput struct {
	_ struct{} `type:"structure"`

	CpuNum *float64 `type:"double"`

	MaxConn *int64 `type:"int64"`

	MaxStorage *int64 `type:"int64"`

	MemInGb *float64 `type:"double"`

	SpecName *string `type:"string"`
}

// String returns the string representation
func (s NodeSpecForDescribeNodeSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeSpecForDescribeNodeSpecsOutput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *NodeSpecForDescribeNodeSpecsOutput) SetCpuNum(v float64) *NodeSpecForDescribeNodeSpecsOutput {
	s.CpuNum = &v
	return s
}

// SetMaxConn sets the MaxConn field's value.
func (s *NodeSpecForDescribeNodeSpecsOutput) SetMaxConn(v int64) *NodeSpecForDescribeNodeSpecsOutput {
	s.MaxConn = &v
	return s
}

// SetMaxStorage sets the MaxStorage field's value.
func (s *NodeSpecForDescribeNodeSpecsOutput) SetMaxStorage(v int64) *NodeSpecForDescribeNodeSpecsOutput {
	s.MaxStorage = &v
	return s
}

// SetMemInGb sets the MemInGb field's value.
func (s *NodeSpecForDescribeNodeSpecsOutput) SetMemInGb(v float64) *NodeSpecForDescribeNodeSpecsOutput {
	s.MemInGb = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *NodeSpecForDescribeNodeSpecsOutput) SetSpecName(v string) *NodeSpecForDescribeNodeSpecsOutput {
	s.SpecName = &v
	return s
}

type ShardNodeSpecForDescribeNodeSpecsOutput struct {
	_ struct{} `type:"structure"`

	CpuNum *float64 `type:"double"`

	MaxConn *int64 `type:"int64"`

	MaxStorage *int64 `type:"int64"`

	MemInGb *float64 `type:"double"`

	SpecName *string `type:"string"`
}

// String returns the string representation
func (s ShardNodeSpecForDescribeNodeSpecsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ShardNodeSpecForDescribeNodeSpecsOutput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *ShardNodeSpecForDescribeNodeSpecsOutput) SetCpuNum(v float64) *ShardNodeSpecForDescribeNodeSpecsOutput {
	s.CpuNum = &v
	return s
}

// SetMaxConn sets the MaxConn field's value.
func (s *ShardNodeSpecForDescribeNodeSpecsOutput) SetMaxConn(v int64) *ShardNodeSpecForDescribeNodeSpecsOutput {
	s.MaxConn = &v
	return s
}

// SetMaxStorage sets the MaxStorage field's value.
func (s *ShardNodeSpecForDescribeNodeSpecsOutput) SetMaxStorage(v int64) *ShardNodeSpecForDescribeNodeSpecsOutput {
	s.MaxStorage = &v
	return s
}

// SetMemInGb sets the MemInGb field's value.
func (s *ShardNodeSpecForDescribeNodeSpecsOutput) SetMemInGb(v float64) *ShardNodeSpecForDescribeNodeSpecsOutput {
	s.MemInGb = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *ShardNodeSpecForDescribeNodeSpecsOutput) SetSpecName(v string) *ShardNodeSpecForDescribeNodeSpecsOutput {
	s.SpecName = &v
	return s
}
