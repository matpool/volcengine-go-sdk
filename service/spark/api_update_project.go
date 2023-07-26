// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateProjectCommon = "updateProject"

// UpdateProjectCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateProjectCommon operation. The "output" return
// value will be populated with the UpdateProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateProjectCommon Send returns without error.
//
// See UpdateProjectCommon for more information on using the UpdateProjectCommon
// API call, and error handling.
//
//	// Example sending a request using the UpdateProjectCommonRequest method.
//	req, resp := client.UpdateProjectCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) UpdateProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateProjectCommon,
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

// UpdateProjectCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation UpdateProjectCommon for usage and error information.
func (c *SPARK) UpdateProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateProjectCommonRequest(input)
	return out, req.Send()
}

// UpdateProjectCommonWithContext is the same as UpdateProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) UpdateProjectCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateProject = "updateProject"

// UpdateProjectRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateProject operation. The "output" return
// value will be populated with the UpdateProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateProjectCommon Send returns without error.
//
// See UpdateProject for more information on using the UpdateProject
// API call, and error handling.
//
//	// Example sending a request using the UpdateProjectRequest method.
//	req, resp := client.UpdateProjectRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) UpdateProjectRequest(input *UpdateProjectInput) (req *request.Request, output *UpdateProjectOutput) {
	op := &request.Operation{
		Name:       opUpdateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateProjectInput{}
	}

	output = &UpdateProjectOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateProject API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation UpdateProject for usage and error information.
func (c *SPARK) UpdateProject(input *UpdateProjectInput) (*UpdateProjectOutput, error) {
	req, out := c.UpdateProjectRequest(input)
	return out, req.Send()
}

// UpdateProjectWithContext is the same as UpdateProject with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) UpdateProjectWithContext(ctx volcengine.Context, input *UpdateProjectInput, opts ...request.Option) (*UpdateProjectOutput, error) {
	req, out := c.UpdateProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateProjectInput struct {
	_ struct{} `type:"structure"`

	AuthorityType *string `type:"string" enum:"EnumOfAuthorityTypeForupdateProjectInput"`

	Description *string `type:"string"`

	OwnerId *string `type:"string"`

	ProjectId *string `type:"string"`

	ProjectName *string `type:"string"`

	ProjectType *string `type:"string" enum:"EnumOfProjectTypeForupdateProjectInput"`
}

// String returns the string representation
func (s UpdateProjectInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateProjectInput) GoString() string {
	return s.String()
}

// SetAuthorityType sets the AuthorityType field's value.
func (s *UpdateProjectInput) SetAuthorityType(v string) *UpdateProjectInput {
	s.AuthorityType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateProjectInput) SetDescription(v string) *UpdateProjectInput {
	s.Description = &v
	return s
}

// SetOwnerId sets the OwnerId field's value.
func (s *UpdateProjectInput) SetOwnerId(v string) *UpdateProjectInput {
	s.OwnerId = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *UpdateProjectInput) SetProjectId(v string) *UpdateProjectInput {
	s.ProjectId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateProjectInput) SetProjectName(v string) *UpdateProjectInput {
	s.ProjectName = &v
	return s
}

// SetProjectType sets the ProjectType field's value.
func (s *UpdateProjectInput) SetProjectType(v string) *UpdateProjectInput {
	s.ProjectType = &v
	return s
}

type UpdateProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Message *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s UpdateProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateProjectOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *UpdateProjectOutput) SetMessage(v string) *UpdateProjectOutput {
	s.Message = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateProjectOutput) SetStatus(v string) *UpdateProjectOutput {
	s.Status = &v
	return s
}

const (
	// EnumOfAuthorityTypeForupdateProjectInputPublic is a EnumOfAuthorityTypeForupdateProjectInput enum value
	EnumOfAuthorityTypeForupdateProjectInputPublic = "PUBLIC"

	// EnumOfAuthorityTypeForupdateProjectInputPrivate is a EnumOfAuthorityTypeForupdateProjectInput enum value
	EnumOfAuthorityTypeForupdateProjectInputPrivate = "PRIVATE"
)

const (
	// EnumOfProjectTypeForupdateProjectInputDefault is a EnumOfProjectTypeForupdateProjectInput enum value
	EnumOfProjectTypeForupdateProjectInputDefault = "DEFAULT"

	// EnumOfProjectTypeForupdateProjectInputBmq is a EnumOfProjectTypeForupdateProjectInput enum value
	EnumOfProjectTypeForupdateProjectInputBmq = "BMQ"

	// EnumOfProjectTypeForupdateProjectInputFlink is a EnumOfProjectTypeForupdateProjectInput enum value
	EnumOfProjectTypeForupdateProjectInputFlink = "FLINK"

	// EnumOfProjectTypeForupdateProjectInputSpark is a EnumOfProjectTypeForupdateProjectInput enum value
	EnumOfProjectTypeForupdateProjectInputSpark = "SPARK"

	// EnumOfProjectTypeForupdateProjectInputCfs is a EnumOfProjectTypeForupdateProjectInput enum value
	EnumOfProjectTypeForupdateProjectInputCfs = "CFS"

	// EnumOfProjectTypeForupdateProjectInputEscloud is a EnumOfProjectTypeForupdateProjectInput enum value
	EnumOfProjectTypeForupdateProjectInputEscloud = "ESCLOUD"
)
