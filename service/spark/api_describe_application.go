// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeApplicationCommon = "describeApplication"

// DescribeApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeApplicationCommon operation. The "output" return
// value will be populated with the DescribeApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeApplicationCommon Send returns without error.
//
// See DescribeApplicationCommon for more information on using the DescribeApplicationCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeApplicationCommonRequest method.
//	req, resp := client.DescribeApplicationCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) DescribeApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeApplicationCommon,
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

// DescribeApplicationCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation DescribeApplicationCommon for usage and error information.
func (c *SPARK) DescribeApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeApplicationCommonRequest(input)
	return out, req.Send()
}

// DescribeApplicationCommonWithContext is the same as DescribeApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) DescribeApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeApplication = "describeApplication"

// DescribeApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeApplication operation. The "output" return
// value will be populated with the DescribeApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeApplicationCommon Send returns without error.
//
// See DescribeApplication for more information on using the DescribeApplication
// API call, and error handling.
//
//	// Example sending a request using the DescribeApplicationRequest method.
//	req, resp := client.DescribeApplicationRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) DescribeApplicationRequest(input *DescribeApplicationInput) (req *request.Request, output *DescribeApplicationOutput) {
	op := &request.Operation{
		Name:       opDescribeApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeApplicationInput{}
	}

	output = &DescribeApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeApplication API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation DescribeApplication for usage and error information.
func (c *SPARK) DescribeApplication(input *DescribeApplicationInput) (*DescribeApplicationOutput, error) {
	req, out := c.DescribeApplicationRequest(input)
	return out, req.Send()
}

// DescribeApplicationWithContext is the same as DescribeApplication with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) DescribeApplicationWithContext(ctx volcengine.Context, input *DescribeApplicationInput, opts ...request.Option) (*DescribeApplicationOutput, error) {
	req, out := c.DescribeApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DependencyFordescribeApplicationOutput struct {
	_ struct{} `type:"structure"`

	Archives []*string `type:"list"`

	Files []*string `type:"list"`

	Jars []*string `type:"list"`

	PyFiles []*string `type:"list"`
}

// String returns the string representation
func (s DependencyFordescribeApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DependencyFordescribeApplicationOutput) GoString() string {
	return s.String()
}

// SetArchives sets the Archives field's value.
func (s *DependencyFordescribeApplicationOutput) SetArchives(v []*string) *DependencyFordescribeApplicationOutput {
	s.Archives = v
	return s
}

// SetFiles sets the Files field's value.
func (s *DependencyFordescribeApplicationOutput) SetFiles(v []*string) *DependencyFordescribeApplicationOutput {
	s.Files = v
	return s
}

// SetJars sets the Jars field's value.
func (s *DependencyFordescribeApplicationOutput) SetJars(v []*string) *DependencyFordescribeApplicationOutput {
	s.Jars = v
	return s
}

// SetPyFiles sets the PyFiles field's value.
func (s *DependencyFordescribeApplicationOutput) SetPyFiles(v []*string) *DependencyFordescribeApplicationOutput {
	s.PyFiles = v
	return s
}

type DeployRequestFordescribeApplicationOutput struct {
	_ struct{} `type:"structure"`

	Priority *string `type:"string"`

	ResourcePoolTrn *string `type:"string"`

	SchedulePolicy *string `type:"string"`

	ScheduleTimeout *string `type:"string"`
}

// String returns the string representation
func (s DeployRequestFordescribeApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeployRequestFordescribeApplicationOutput) GoString() string {
	return s.String()
}

// SetPriority sets the Priority field's value.
func (s *DeployRequestFordescribeApplicationOutput) SetPriority(v string) *DeployRequestFordescribeApplicationOutput {
	s.Priority = &v
	return s
}

// SetResourcePoolTrn sets the ResourcePoolTrn field's value.
func (s *DeployRequestFordescribeApplicationOutput) SetResourcePoolTrn(v string) *DeployRequestFordescribeApplicationOutput {
	s.ResourcePoolTrn = &v
	return s
}

// SetSchedulePolicy sets the SchedulePolicy field's value.
func (s *DeployRequestFordescribeApplicationOutput) SetSchedulePolicy(v string) *DeployRequestFordescribeApplicationOutput {
	s.SchedulePolicy = &v
	return s
}

// SetScheduleTimeout sets the ScheduleTimeout field's value.
func (s *DeployRequestFordescribeApplicationOutput) SetScheduleTimeout(v string) *DeployRequestFordescribeApplicationOutput {
	s.ScheduleTimeout = &v
	return s
}

type DescribeApplicationInput struct {
	_ struct{} `type:"structure"`

	ApplicationTrn *string `type:"string"`
}

// String returns the string representation
func (s DescribeApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeApplicationInput) GoString() string {
	return s.String()
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *DescribeApplicationInput) SetApplicationTrn(v string) *DescribeApplicationInput {
	s.ApplicationTrn = &v
	return s
}

type DescribeApplicationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Alert *bool `type:"boolean"`

	ApplicationName *string `type:"string"`

	ApplicationTrn *string `type:"string"`

	ApplicationType *string `type:"string"`

	Args *string `type:"string"`

	Conf map[string]*string `type:"map"`

	Dependency *DependencyFordescribeApplicationOutput `type:"structure"`

	DeployRequest *DeployRequestFordescribeApplicationOutput `type:"structure"`

	EngineVersion *string `type:"string"`

	Image *string `type:"string"`

	IsLatestVersion *bool `type:"boolean"`

	Jar *string `type:"string"`

	LatestVersion *string `type:"string"`

	MainClass *string `type:"string"`

	ProjectId *string `type:"string"`

	RestUrl *string `type:"string"`

	SqlText *string `type:"string"`

	State *string `type:"string"`

	UserId *string `type:"string"`

	VersionName *string `type:"string"`
}

// String returns the string representation
func (s DescribeApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeApplicationOutput) GoString() string {
	return s.String()
}

// SetAlert sets the Alert field's value.
func (s *DescribeApplicationOutput) SetAlert(v bool) *DescribeApplicationOutput {
	s.Alert = &v
	return s
}

// SetApplicationName sets the ApplicationName field's value.
func (s *DescribeApplicationOutput) SetApplicationName(v string) *DescribeApplicationOutput {
	s.ApplicationName = &v
	return s
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *DescribeApplicationOutput) SetApplicationTrn(v string) *DescribeApplicationOutput {
	s.ApplicationTrn = &v
	return s
}

// SetApplicationType sets the ApplicationType field's value.
func (s *DescribeApplicationOutput) SetApplicationType(v string) *DescribeApplicationOutput {
	s.ApplicationType = &v
	return s
}

// SetArgs sets the Args field's value.
func (s *DescribeApplicationOutput) SetArgs(v string) *DescribeApplicationOutput {
	s.Args = &v
	return s
}

// SetConf sets the Conf field's value.
func (s *DescribeApplicationOutput) SetConf(v map[string]*string) *DescribeApplicationOutput {
	s.Conf = v
	return s
}

// SetDependency sets the Dependency field's value.
func (s *DescribeApplicationOutput) SetDependency(v *DependencyFordescribeApplicationOutput) *DescribeApplicationOutput {
	s.Dependency = v
	return s
}

// SetDeployRequest sets the DeployRequest field's value.
func (s *DescribeApplicationOutput) SetDeployRequest(v *DeployRequestFordescribeApplicationOutput) *DescribeApplicationOutput {
	s.DeployRequest = v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *DescribeApplicationOutput) SetEngineVersion(v string) *DescribeApplicationOutput {
	s.EngineVersion = &v
	return s
}

// SetImage sets the Image field's value.
func (s *DescribeApplicationOutput) SetImage(v string) *DescribeApplicationOutput {
	s.Image = &v
	return s
}

// SetIsLatestVersion sets the IsLatestVersion field's value.
func (s *DescribeApplicationOutput) SetIsLatestVersion(v bool) *DescribeApplicationOutput {
	s.IsLatestVersion = &v
	return s
}

// SetJar sets the Jar field's value.
func (s *DescribeApplicationOutput) SetJar(v string) *DescribeApplicationOutput {
	s.Jar = &v
	return s
}

// SetLatestVersion sets the LatestVersion field's value.
func (s *DescribeApplicationOutput) SetLatestVersion(v string) *DescribeApplicationOutput {
	s.LatestVersion = &v
	return s
}

// SetMainClass sets the MainClass field's value.
func (s *DescribeApplicationOutput) SetMainClass(v string) *DescribeApplicationOutput {
	s.MainClass = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *DescribeApplicationOutput) SetProjectId(v string) *DescribeApplicationOutput {
	s.ProjectId = &v
	return s
}

// SetRestUrl sets the RestUrl field's value.
func (s *DescribeApplicationOutput) SetRestUrl(v string) *DescribeApplicationOutput {
	s.RestUrl = &v
	return s
}

// SetSqlText sets the SqlText field's value.
func (s *DescribeApplicationOutput) SetSqlText(v string) *DescribeApplicationOutput {
	s.SqlText = &v
	return s
}

// SetState sets the State field's value.
func (s *DescribeApplicationOutput) SetState(v string) *DescribeApplicationOutput {
	s.State = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *DescribeApplicationOutput) SetUserId(v string) *DescribeApplicationOutput {
	s.UserId = &v
	return s
}

// SetVersionName sets the VersionName field's value.
func (s *DescribeApplicationOutput) SetVersionName(v string) *DescribeApplicationOutput {
	s.VersionName = &v
	return s
}
