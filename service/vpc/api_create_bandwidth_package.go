// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateBandwidthPackageCommon = "CreateBandwidthPackage"

// CreateBandwidthPackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateBandwidthPackageCommon operation. The "output" return
// value will be populated with the CreateBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateBandwidthPackageCommon Send returns without error.
//
// See CreateBandwidthPackageCommon for more information on using the CreateBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateBandwidthPackageCommonRequest method.
//    req, resp := client.CreateBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateBandwidthPackageCommon,
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

// CreateBandwidthPackageCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateBandwidthPackageCommon for usage and error information.
func (c *VPC) CreateBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// CreateBandwidthPackageCommonWithContext is the same as CreateBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateBandwidthPackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateBandwidthPackage = "CreateBandwidthPackage"

// CreateBandwidthPackageRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateBandwidthPackage operation. The "output" return
// value will be populated with the CreateBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateBandwidthPackageCommon Send returns without error.
//
// See CreateBandwidthPackage for more information on using the CreateBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the CreateBandwidthPackageRequest method.
//    req, resp := client.CreateBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateBandwidthPackageRequest(input *CreateBandwidthPackageInput) (req *request.Request, output *CreateBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opCreateBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateBandwidthPackageInput{}
	}

	output = &CreateBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateBandwidthPackage API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateBandwidthPackage for usage and error information.
func (c *VPC) CreateBandwidthPackage(input *CreateBandwidthPackageInput) (*CreateBandwidthPackageOutput, error) {
	req, out := c.CreateBandwidthPackageRequest(input)
	return out, req.Send()
}

// CreateBandwidthPackageWithContext is the same as CreateBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See CreateBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateBandwidthPackageWithContext(ctx volcengine.Context, input *CreateBandwidthPackageInput, opts ...request.Option) (*CreateBandwidthPackageOutput, error) {
	req, out := c.CreateBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	// Bandwidth is a required field
	Bandwidth *int64 `min:"2" max:"5000" type:"integer" required:"true"`

	BandwidthPackageName *string `min:"1" max:"128" type:"string"`

	BillingType *int64 `min:"2" max:"4" type:"integer"`

	Description *string `min:"1" max:"255" type:"string"`

	ISP *string `type:"string" enum:"ISPForCreateBandwidthPackageInput"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateBandwidthPackageInput `type:"list"`
}

// String returns the string representation
func (s CreateBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateBandwidthPackageInput"}
	if s.Bandwidth == nil {
		invalidParams.Add(request.NewErrParamRequired("Bandwidth"))
	}
	if s.Bandwidth != nil && *s.Bandwidth < 2 {
		invalidParams.Add(request.NewErrParamMinValue("Bandwidth", 2))
	}
	if s.Bandwidth != nil && *s.Bandwidth > 5000 {
		invalidParams.Add(request.NewErrParamMaxValue("Bandwidth", 5000))
	}
	if s.BandwidthPackageName != nil && len(*s.BandwidthPackageName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("BandwidthPackageName", 1))
	}
	if s.BandwidthPackageName != nil && len(*s.BandwidthPackageName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("BandwidthPackageName", 128, *s.BandwidthPackageName))
	}
	if s.BillingType != nil && *s.BillingType < 2 {
		invalidParams.Add(request.NewErrParamMinValue("BillingType", 2))
	}
	if s.BillingType != nil && *s.BillingType > 4 {
		invalidParams.Add(request.NewErrParamMaxValue("BillingType", 4))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CreateBandwidthPackageInput) SetBandwidth(v int64) *CreateBandwidthPackageInput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageName sets the BandwidthPackageName field's value.
func (s *CreateBandwidthPackageInput) SetBandwidthPackageName(v string) *CreateBandwidthPackageInput {
	s.BandwidthPackageName = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *CreateBandwidthPackageInput) SetBillingType(v int64) *CreateBandwidthPackageInput {
	s.BillingType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateBandwidthPackageInput) SetDescription(v string) *CreateBandwidthPackageInput {
	s.Description = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *CreateBandwidthPackageInput) SetISP(v string) *CreateBandwidthPackageInput {
	s.ISP = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateBandwidthPackageInput) SetProjectName(v string) *CreateBandwidthPackageInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateBandwidthPackageInput) SetTags(v []*TagForCreateBandwidthPackageInput) *CreateBandwidthPackageInput {
	s.Tags = v
	return s
}

type CreateBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BandwidthPackageId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateBandwidthPackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBandwidthPackageOutput) GoString() string {
	return s.String()
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *CreateBandwidthPackageOutput) SetBandwidthPackageId(v string) *CreateBandwidthPackageOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateBandwidthPackageOutput) SetRequestId(v string) *CreateBandwidthPackageOutput {
	s.RequestId = &v
	return s
}

type TagForCreateBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateBandwidthPackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateBandwidthPackageInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateBandwidthPackageInput) SetKey(v string) *TagForCreateBandwidthPackageInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateBandwidthPackageInput) SetValue(v string) *TagForCreateBandwidthPackageInput {
	s.Value = &v
	return s
}

const (
	// ISPForCreateBandwidthPackageInputBgp is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputBgp = "BGP"

	// ISPForCreateBandwidthPackageInputChinaMobile is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaMobile = "ChinaMobile"

	// ISPForCreateBandwidthPackageInputChinaUnicom is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaUnicom = "ChinaUnicom"

	// ISPForCreateBandwidthPackageInputChinaTelecom is a ISPForCreateBandwidthPackageInput enum value
	ISPForCreateBandwidthPackageInputChinaTelecom = "ChinaTelecom"
)
