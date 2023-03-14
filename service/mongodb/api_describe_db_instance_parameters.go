// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstanceParametersCommon = "DescribeDBInstanceParameters"

// DescribeDBInstanceParametersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceParametersCommon operation. The "output" return
// value will be populated with the DescribeDBInstanceParametersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceParametersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceParametersCommon Send returns without error.
//
// See DescribeDBInstanceParametersCommon for more information on using the DescribeDBInstanceParametersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceParametersCommonRequest method.
//    req, resp := client.DescribeDBInstanceParametersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeDBInstanceParametersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceParametersCommon,
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

// DescribeDBInstanceParametersCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBInstanceParametersCommon for usage and error information.
func (c *MONGODB) DescribeDBInstanceParametersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceParametersCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceParametersCommonWithContext is the same as DescribeDBInstanceParametersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceParametersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeDBInstanceParametersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceParametersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstanceParameters = "DescribeDBInstanceParameters"

// DescribeDBInstanceParametersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceParameters operation. The "output" return
// value will be populated with the DescribeDBInstanceParametersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceParametersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceParametersCommon Send returns without error.
//
// See DescribeDBInstanceParameters for more information on using the DescribeDBInstanceParameters
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceParametersRequest method.
//    req, resp := client.DescribeDBInstanceParametersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeDBInstanceParametersRequest(input *DescribeDBInstanceParametersInput) (req *request.Request, output *DescribeDBInstanceParametersOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstanceParametersInput{}
	}

	output = &DescribeDBInstanceParametersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstanceParameters API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBInstanceParameters for usage and error information.
func (c *MONGODB) DescribeDBInstanceParameters(input *DescribeDBInstanceParametersInput) (*DescribeDBInstanceParametersOutput, error) {
	req, out := c.DescribeDBInstanceParametersRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceParametersWithContext is the same as DescribeDBInstanceParameters with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceParameters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeDBInstanceParametersWithContext(ctx volcengine.Context, input *DescribeDBInstanceParametersInput, opts ...request.Option) (*DescribeDBInstanceParametersOutput, error) {
	req, out := c.DescribeDBInstanceParametersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDBInstanceParametersInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	ParameterNames *string `type:"string"`

	ParameterRole *string `type:"string" enum:"EnumOfParameterRoleForDescribeDBInstanceParametersInput"`
}

// String returns the string representation
func (s DescribeDBInstanceParametersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceParametersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceParametersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstanceParametersInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceParametersInput) SetInstanceId(v string) *DescribeDBInstanceParametersInput {
	s.InstanceId = &v
	return s
}

// SetParameterNames sets the ParameterNames field's value.
func (s *DescribeDBInstanceParametersInput) SetParameterNames(v string) *DescribeDBInstanceParametersInput {
	s.ParameterNames = &v
	return s
}

// SetParameterRole sets the ParameterRole field's value.
func (s *DescribeDBInstanceParametersInput) SetParameterRole(v string) *DescribeDBInstanceParametersInput {
	s.ParameterRole = &v
	return s
}

type DescribeDBInstanceParametersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DBEngine *string `type:"string"`

	DBEngineVersion *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceParameters []*InstanceParameterForDescribeDBInstanceParametersOutput `type:"list"`

	Total *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBInstanceParametersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceParametersOutput) GoString() string {
	return s.String()
}

// SetDBEngine sets the DBEngine field's value.
func (s *DescribeDBInstanceParametersOutput) SetDBEngine(v string) *DescribeDBInstanceParametersOutput {
	s.DBEngine = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *DescribeDBInstanceParametersOutput) SetDBEngineVersion(v string) *DescribeDBInstanceParametersOutput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceParametersOutput) SetInstanceId(v string) *DescribeDBInstanceParametersOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceParameters sets the InstanceParameters field's value.
func (s *DescribeDBInstanceParametersOutput) SetInstanceParameters(v []*InstanceParameterForDescribeDBInstanceParametersOutput) *DescribeDBInstanceParametersOutput {
	s.InstanceParameters = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBInstanceParametersOutput) SetTotal(v string) *DescribeDBInstanceParametersOutput {
	s.Total = &v
	return s
}

type InstanceParameterForDescribeDBInstanceParametersOutput struct {
	_ struct{} `type:"structure"`

	CheckingCode *string `type:"string"`

	ForceModify *bool `type:"boolean"`

	ForceRestart *bool `type:"boolean"`

	ParameterDefaultValue *string `type:"string"`

	ParameterDescription *string `type:"string"`

	ParameterRole *string `type:"string"`

	ParameterType *string `type:"string"`

	ParameterValue *string `type:"string"`
}

// String returns the string representation
func (s InstanceParameterForDescribeDBInstanceParametersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceParameterForDescribeDBInstanceParametersOutput) GoString() string {
	return s.String()
}

// SetCheckingCode sets the CheckingCode field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetCheckingCode(v string) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.CheckingCode = &v
	return s
}

// SetForceModify sets the ForceModify field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetForceModify(v bool) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ForceModify = &v
	return s
}

// SetForceRestart sets the ForceRestart field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetForceRestart(v bool) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ForceRestart = &v
	return s
}

// SetParameterDefaultValue sets the ParameterDefaultValue field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetParameterDefaultValue(v string) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ParameterDefaultValue = &v
	return s
}

// SetParameterDescription sets the ParameterDescription field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetParameterDescription(v string) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ParameterDescription = &v
	return s
}

// SetParameterRole sets the ParameterRole field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetParameterRole(v string) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ParameterRole = &v
	return s
}

// SetParameterType sets the ParameterType field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetParameterType(v string) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ParameterType = &v
	return s
}

// SetParameterValue sets the ParameterValue field's value.
func (s *InstanceParameterForDescribeDBInstanceParametersOutput) SetParameterValue(v string) *InstanceParameterForDescribeDBInstanceParametersOutput {
	s.ParameterValue = &v
	return s
}

const (
	// EnumOfParameterRoleForDescribeDBInstanceParametersInputConfigServer is a EnumOfParameterRoleForDescribeDBInstanceParametersInput enum value
	EnumOfParameterRoleForDescribeDBInstanceParametersInputConfigServer = "ConfigServer"

	// EnumOfParameterRoleForDescribeDBInstanceParametersInputMongos is a EnumOfParameterRoleForDescribeDBInstanceParametersInput enum value
	EnumOfParameterRoleForDescribeDBInstanceParametersInputMongos = "Mongos"

	// EnumOfParameterRoleForDescribeDBInstanceParametersInputNode is a EnumOfParameterRoleForDescribeDBInstanceParametersInput enum value
	EnumOfParameterRoleForDescribeDBInstanceParametersInputNode = "Node"

	// EnumOfParameterRoleForDescribeDBInstanceParametersInputShard is a EnumOfParameterRoleForDescribeDBInstanceParametersInput enum value
	EnumOfParameterRoleForDescribeDBInstanceParametersInputShard = "Shard"

	// EnumOfParameterRoleForDescribeDBInstanceParametersInputUnknown is a EnumOfParameterRoleForDescribeDBInstanceParametersInput enum value
	EnumOfParameterRoleForDescribeDBInstanceParametersInputUnknown = "Unknown"
)
