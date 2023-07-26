// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCopyImageCommon = "CopyImage"

// CopyImageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CopyImageCommon operation. The "output" return
// value will be populated with the CopyImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CopyImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CopyImageCommon Send returns without error.
//
// See CopyImageCommon for more information on using the CopyImageCommon
// API call, and error handling.
//
//	// Example sending a request using the CopyImageCommonRequest method.
//	req, resp := client.CopyImageCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) CopyImageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCopyImageCommon,
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

// CopyImageCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CopyImageCommon for usage and error information.
func (c *ECS) CopyImageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CopyImageCommonRequest(input)
	return out, req.Send()
}

// CopyImageCommonWithContext is the same as CopyImageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CopyImageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CopyImageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CopyImageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCopyImage = "CopyImage"

// CopyImageRequest generates a "volcengine/request.Request" representing the
// client's request for the CopyImage operation. The "output" return
// value will be populated with the CopyImageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CopyImageCommon Request to send the API call to the service.
// the "output" return value is not valid until after CopyImageCommon Send returns without error.
//
// See CopyImage for more information on using the CopyImage
// API call, and error handling.
//
//	// Example sending a request using the CopyImageRequest method.
//	req, resp := client.CopyImageRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) CopyImageRequest(input *CopyImageInput) (req *request.Request, output *CopyImageOutput) {
	op := &request.Operation{
		Name:       opCopyImage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyImageInput{}
	}

	output = &CopyImageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CopyImage API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CopyImage for usage and error information.
func (c *ECS) CopyImage(input *CopyImageInput) (*CopyImageOutput, error) {
	req, out := c.CopyImageRequest(input)
	return out, req.Send()
}

// CopyImageWithContext is the same as CopyImage with the addition of
// the ability to pass a context and additional request options.
//
// See CopyImage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CopyImageWithContext(ctx volcengine.Context, input *CopyImageInput, opts ...request.Option) (*CopyImageOutput, error) {
	req, out := c.CopyImageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CopyImageInput struct {
	_ struct{} `type:"structure"`

	CopyImageTags *bool `type:"boolean"`

	Description *string `type:"string"`

	DestinationRegion *string `type:"string"`

	ImageId *string `type:"string"`

	ImageName *string `type:"string"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s CopyImageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CopyImageInput) GoString() string {
	return s.String()
}

// SetCopyImageTags sets the CopyImageTags field's value.
func (s *CopyImageInput) SetCopyImageTags(v bool) *CopyImageInput {
	s.CopyImageTags = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CopyImageInput) SetDescription(v string) *CopyImageInput {
	s.Description = &v
	return s
}

// SetDestinationRegion sets the DestinationRegion field's value.
func (s *CopyImageInput) SetDestinationRegion(v string) *CopyImageInput {
	s.DestinationRegion = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *CopyImageInput) SetImageId(v string) *CopyImageInput {
	s.ImageId = &v
	return s
}

// SetImageName sets the ImageName field's value.
func (s *CopyImageInput) SetImageName(v string) *CopyImageInput {
	s.ImageName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CopyImageInput) SetProjectName(v string) *CopyImageInput {
	s.ProjectName = &v
	return s
}

type CopyImageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ImageId *string `type:"string"`
}

// String returns the string representation
func (s CopyImageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CopyImageOutput) GoString() string {
	return s.String()
}

// SetImageId sets the ImageId field's value.
func (s *CopyImageOutput) SetImageId(v string) *CopyImageOutput {
	s.ImageId = &v
	return s
}
