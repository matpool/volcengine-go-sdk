// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateResourcePoolCommon = "createResourcePool"

// CreateResourcePoolCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateResourcePoolCommon operation. The "output" return
// value will be populated with the CreateResourcePoolCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateResourcePoolCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateResourcePoolCommon Send returns without error.
//
// See CreateResourcePoolCommon for more information on using the CreateResourcePoolCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateResourcePoolCommonRequest method.
//	req, resp := client.CreateResourcePoolCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) CreateResourcePoolCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateResourcePoolCommon,
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

// CreateResourcePoolCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateResourcePoolCommon for usage and error information.
func (c *SPARK) CreateResourcePoolCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateResourcePoolCommonRequest(input)
	return out, req.Send()
}

// CreateResourcePoolCommonWithContext is the same as CreateResourcePoolCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourcePoolCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateResourcePoolCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateResourcePoolCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateResourcePool = "createResourcePool"

// CreateResourcePoolRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateResourcePool operation. The "output" return
// value will be populated with the CreateResourcePoolCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateResourcePoolCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateResourcePoolCommon Send returns without error.
//
// See CreateResourcePool for more information on using the CreateResourcePool
// API call, and error handling.
//
//	// Example sending a request using the CreateResourcePoolRequest method.
//	req, resp := client.CreateResourcePoolRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) CreateResourcePoolRequest(input *CreateResourcePoolInput) (req *request.Request, output *CreateResourcePoolOutput) {
	op := &request.Operation{
		Name:       opCreateResourcePool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResourcePoolInput{}
	}

	output = &CreateResourcePoolOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateResourcePool API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateResourcePool for usage and error information.
func (c *SPARK) CreateResourcePool(input *CreateResourcePoolInput) (*CreateResourcePoolOutput, error) {
	req, out := c.CreateResourcePoolRequest(input)
	return out, req.Send()
}

// CreateResourcePoolWithContext is the same as CreateResourcePool with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourcePool for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateResourcePoolWithContext(ctx volcengine.Context, input *CreateResourcePoolInput, opts ...request.Option) (*CreateResourcePoolOutput, error) {
	req, out := c.CreateResourcePoolRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateResourcePoolInput struct {
	_ struct{} `type:"structure"`

	// BillingType is a required field
	BillingType *string `type:"string" required:"true"`

	// Name is a required field
	Name *string `max:"30" type:"string" required:"true"`

	// ProjectId is a required field
	ProjectId *string `type:"string" required:"true"`

	Resources []*ResourceForcreateResourcePoolInput `type:"list"`

	SecurityGroupIdList []*string `type:"list"`

	SubnetIdList []*string `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	// ZoneId is a required field
	ZoneId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResourcePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateResourcePoolInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourcePoolInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateResourcePoolInput"}
	if s.BillingType == nil {
		invalidParams.Add(request.NewErrParamRequired("BillingType"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) > 30 {
		invalidParams.Add(request.NewErrParamMaxLen("Name", 30, *s.Name))
	}
	if s.ProjectId == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}
	if s.ZoneId == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBillingType sets the BillingType field's value.
func (s *CreateResourcePoolInput) SetBillingType(v string) *CreateResourcePoolInput {
	s.BillingType = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateResourcePoolInput) SetName(v string) *CreateResourcePoolInput {
	s.Name = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *CreateResourcePoolInput) SetProjectId(v string) *CreateResourcePoolInput {
	s.ProjectId = &v
	return s
}

// SetResources sets the Resources field's value.
func (s *CreateResourcePoolInput) SetResources(v []*ResourceForcreateResourcePoolInput) *CreateResourcePoolInput {
	s.Resources = v
	return s
}

// SetSecurityGroupIdList sets the SecurityGroupIdList field's value.
func (s *CreateResourcePoolInput) SetSecurityGroupIdList(v []*string) *CreateResourcePoolInput {
	s.SecurityGroupIdList = v
	return s
}

// SetSubnetIdList sets the SubnetIdList field's value.
func (s *CreateResourcePoolInput) SetSubnetIdList(v []*string) *CreateResourcePoolInput {
	s.SubnetIdList = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateResourcePoolInput) SetVpcId(v string) *CreateResourcePoolInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateResourcePoolInput) SetZoneId(v string) *CreateResourcePoolInput {
	s.ZoneId = &v
	return s
}

type CreateResourcePoolOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResourcePoolTrn *string `type:"string"`
}

// String returns the string representation
func (s CreateResourcePoolOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateResourcePoolOutput) GoString() string {
	return s.String()
}

// SetResourcePoolTrn sets the ResourcePoolTrn field's value.
func (s *CreateResourcePoolOutput) SetResourcePoolTrn(v string) *CreateResourcePoolOutput {
	s.ResourcePoolTrn = &v
	return s
}

type ResourceForcreateResourcePoolInput struct {
	_ struct{} `type:"structure"`

	Basic *int64 `type:"int64"`

	Capability *int64 `type:"int64"`

	Units *string `type:"string"`

	Uri *string `type:"string"`
}

// String returns the string representation
func (s ResourceForcreateResourcePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceForcreateResourcePoolInput) GoString() string {
	return s.String()
}

// SetBasic sets the Basic field's value.
func (s *ResourceForcreateResourcePoolInput) SetBasic(v int64) *ResourceForcreateResourcePoolInput {
	s.Basic = &v
	return s
}

// SetCapability sets the Capability field's value.
func (s *ResourceForcreateResourcePoolInput) SetCapability(v int64) *ResourceForcreateResourcePoolInput {
	s.Capability = &v
	return s
}

// SetUnits sets the Units field's value.
func (s *ResourceForcreateResourcePoolInput) SetUnits(v string) *ResourceForcreateResourcePoolInput {
	s.Units = &v
	return s
}

// SetUri sets the Uri field's value.
func (s *ResourceForcreateResourcePoolInput) SetUri(v string) *ResourceForcreateResourcePoolInput {
	s.Uri = &v
	return s
}
