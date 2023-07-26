// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListApplicationCommon = "listApplication"

// ListApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListApplicationCommon operation. The "output" return
// value will be populated with the ListApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListApplicationCommon Send returns without error.
//
// See ListApplicationCommon for more information on using the ListApplicationCommon
// API call, and error handling.
//
//	// Example sending a request using the ListApplicationCommonRequest method.
//	req, resp := client.ListApplicationCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) ListApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListApplicationCommon,
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

// ListApplicationCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListApplicationCommon for usage and error information.
func (c *SPARK) ListApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListApplicationCommonRequest(input)
	return out, req.Send()
}

// ListApplicationCommonWithContext is the same as ListApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListApplication = "listApplication"

// ListApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the ListApplication operation. The "output" return
// value will be populated with the ListApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListApplicationCommon Send returns without error.
//
// See ListApplication for more information on using the ListApplication
// API call, and error handling.
//
//	// Example sending a request using the ListApplicationRequest method.
//	req, resp := client.ListApplicationRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) ListApplicationRequest(input *ListApplicationInput) (req *request.Request, output *ListApplicationOutput) {
	op := &request.Operation{
		Name:       opListApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListApplicationInput{}
	}

	output = &ListApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListApplication API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListApplication for usage and error information.
func (c *SPARK) ListApplication(input *ListApplicationInput) (*ListApplicationOutput, error) {
	req, out := c.ListApplicationRequest(input)
	return out, req.Send()
}

// ListApplicationWithContext is the same as ListApplication with the addition of
// the ability to pass a context and additional request options.
//
// See ListApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListApplicationWithContext(ctx volcengine.Context, input *ListApplicationInput, opts ...request.Option) (*ListApplicationOutput, error) {
	req, out := c.ListApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DependencyForlistApplicationOutput struct {
	_ struct{} `type:"structure"`

	Archives []*string `type:"list"`

	Files []*string `type:"list"`

	Jars []*string `type:"list"`

	PyFiles []*string `type:"list"`
}

// String returns the string representation
func (s DependencyForlistApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DependencyForlistApplicationOutput) GoString() string {
	return s.String()
}

// SetArchives sets the Archives field's value.
func (s *DependencyForlistApplicationOutput) SetArchives(v []*string) *DependencyForlistApplicationOutput {
	s.Archives = v
	return s
}

// SetFiles sets the Files field's value.
func (s *DependencyForlistApplicationOutput) SetFiles(v []*string) *DependencyForlistApplicationOutput {
	s.Files = v
	return s
}

// SetJars sets the Jars field's value.
func (s *DependencyForlistApplicationOutput) SetJars(v []*string) *DependencyForlistApplicationOutput {
	s.Jars = v
	return s
}

// SetPyFiles sets the PyFiles field's value.
func (s *DependencyForlistApplicationOutput) SetPyFiles(v []*string) *DependencyForlistApplicationOutput {
	s.PyFiles = v
	return s
}

type DeployRequestForlistApplicationOutput struct {
	_ struct{} `type:"structure"`

	Priority *string `type:"string"`

	ResourcePoolTrn *string `type:"string"`

	SchedulePolicy *string `type:"string"`

	ScheduleTimeout *string `type:"string"`
}

// String returns the string representation
func (s DeployRequestForlistApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeployRequestForlistApplicationOutput) GoString() string {
	return s.String()
}

// SetPriority sets the Priority field's value.
func (s *DeployRequestForlistApplicationOutput) SetPriority(v string) *DeployRequestForlistApplicationOutput {
	s.Priority = &v
	return s
}

// SetResourcePoolTrn sets the ResourcePoolTrn field's value.
func (s *DeployRequestForlistApplicationOutput) SetResourcePoolTrn(v string) *DeployRequestForlistApplicationOutput {
	s.ResourcePoolTrn = &v
	return s
}

// SetSchedulePolicy sets the SchedulePolicy field's value.
func (s *DeployRequestForlistApplicationOutput) SetSchedulePolicy(v string) *DeployRequestForlistApplicationOutput {
	s.SchedulePolicy = &v
	return s
}

// SetScheduleTimeout sets the ScheduleTimeout field's value.
func (s *DeployRequestForlistApplicationOutput) SetScheduleTimeout(v string) *DeployRequestForlistApplicationOutput {
	s.ScheduleTimeout = &v
	return s
}

type ListApplicationInput struct {
	_ struct{} `type:"structure"`

	ApplicationName *string `type:"string"`

	ApplicationTrn *int64 `type:"int64"`

	PageNum *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectId *string `type:"string"`

	SortFiled *string `type:"string"`

	SortOrder *string `type:"string"`

	State *string `type:"string" enum:"EnumOfStateForlistApplicationInput"`
}

// String returns the string representation
func (s ListApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListApplicationInput) GoString() string {
	return s.String()
}

// SetApplicationName sets the ApplicationName field's value.
func (s *ListApplicationInput) SetApplicationName(v string) *ListApplicationInput {
	s.ApplicationName = &v
	return s
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *ListApplicationInput) SetApplicationTrn(v int64) *ListApplicationInput {
	s.ApplicationTrn = &v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListApplicationInput) SetPageNum(v int32) *ListApplicationInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListApplicationInput) SetPageSize(v int32) *ListApplicationInput {
	s.PageSize = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *ListApplicationInput) SetProjectId(v string) *ListApplicationInput {
	s.ProjectId = &v
	return s
}

// SetSortFiled sets the SortFiled field's value.
func (s *ListApplicationInput) SetSortFiled(v string) *ListApplicationInput {
	s.SortFiled = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListApplicationInput) SetSortOrder(v string) *ListApplicationInput {
	s.SortOrder = &v
	return s
}

// SetState sets the State field's value.
func (s *ListApplicationInput) SetState(v string) *ListApplicationInput {
	s.State = &v
	return s
}

type ListApplicationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Current *string `type:"string"`

	PageSize *string `type:"string"`

	Records []*RecordForlistApplicationOutput `type:"list"`

	Total *string `type:"string"`
}

// String returns the string representation
func (s ListApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListApplicationOutput) GoString() string {
	return s.String()
}

// SetCurrent sets the Current field's value.
func (s *ListApplicationOutput) SetCurrent(v string) *ListApplicationOutput {
	s.Current = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListApplicationOutput) SetPageSize(v string) *ListApplicationOutput {
	s.PageSize = &v
	return s
}

// SetRecords sets the Records field's value.
func (s *ListApplicationOutput) SetRecords(v []*RecordForlistApplicationOutput) *ListApplicationOutput {
	s.Records = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListApplicationOutput) SetTotal(v string) *ListApplicationOutput {
	s.Total = &v
	return s
}

type RecordForlistApplicationOutput struct {
	_ struct{} `type:"structure"`

	Alert *bool `type:"boolean"`

	ApplicationName *string `type:"string"`

	ApplicationTrn *string `type:"string"`

	ApplicationType *string `type:"string"`

	Args *string `type:"string"`

	Conf map[string]*string `type:"map"`

	Dependency *DependencyForlistApplicationOutput `type:"structure"`

	DeployRequest *DeployRequestForlistApplicationOutput `type:"structure"`

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
func (s RecordForlistApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RecordForlistApplicationOutput) GoString() string {
	return s.String()
}

// SetAlert sets the Alert field's value.
func (s *RecordForlistApplicationOutput) SetAlert(v bool) *RecordForlistApplicationOutput {
	s.Alert = &v
	return s
}

// SetApplicationName sets the ApplicationName field's value.
func (s *RecordForlistApplicationOutput) SetApplicationName(v string) *RecordForlistApplicationOutput {
	s.ApplicationName = &v
	return s
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *RecordForlistApplicationOutput) SetApplicationTrn(v string) *RecordForlistApplicationOutput {
	s.ApplicationTrn = &v
	return s
}

// SetApplicationType sets the ApplicationType field's value.
func (s *RecordForlistApplicationOutput) SetApplicationType(v string) *RecordForlistApplicationOutput {
	s.ApplicationType = &v
	return s
}

// SetArgs sets the Args field's value.
func (s *RecordForlistApplicationOutput) SetArgs(v string) *RecordForlistApplicationOutput {
	s.Args = &v
	return s
}

// SetConf sets the Conf field's value.
func (s *RecordForlistApplicationOutput) SetConf(v map[string]*string) *RecordForlistApplicationOutput {
	s.Conf = v
	return s
}

// SetDependency sets the Dependency field's value.
func (s *RecordForlistApplicationOutput) SetDependency(v *DependencyForlistApplicationOutput) *RecordForlistApplicationOutput {
	s.Dependency = v
	return s
}

// SetDeployRequest sets the DeployRequest field's value.
func (s *RecordForlistApplicationOutput) SetDeployRequest(v *DeployRequestForlistApplicationOutput) *RecordForlistApplicationOutput {
	s.DeployRequest = v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *RecordForlistApplicationOutput) SetEngineVersion(v string) *RecordForlistApplicationOutput {
	s.EngineVersion = &v
	return s
}

// SetImage sets the Image field's value.
func (s *RecordForlistApplicationOutput) SetImage(v string) *RecordForlistApplicationOutput {
	s.Image = &v
	return s
}

// SetIsLatestVersion sets the IsLatestVersion field's value.
func (s *RecordForlistApplicationOutput) SetIsLatestVersion(v bool) *RecordForlistApplicationOutput {
	s.IsLatestVersion = &v
	return s
}

// SetJar sets the Jar field's value.
func (s *RecordForlistApplicationOutput) SetJar(v string) *RecordForlistApplicationOutput {
	s.Jar = &v
	return s
}

// SetLatestVersion sets the LatestVersion field's value.
func (s *RecordForlistApplicationOutput) SetLatestVersion(v string) *RecordForlistApplicationOutput {
	s.LatestVersion = &v
	return s
}

// SetMainClass sets the MainClass field's value.
func (s *RecordForlistApplicationOutput) SetMainClass(v string) *RecordForlistApplicationOutput {
	s.MainClass = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *RecordForlistApplicationOutput) SetProjectId(v string) *RecordForlistApplicationOutput {
	s.ProjectId = &v
	return s
}

// SetRestUrl sets the RestUrl field's value.
func (s *RecordForlistApplicationOutput) SetRestUrl(v string) *RecordForlistApplicationOutput {
	s.RestUrl = &v
	return s
}

// SetSqlText sets the SqlText field's value.
func (s *RecordForlistApplicationOutput) SetSqlText(v string) *RecordForlistApplicationOutput {
	s.SqlText = &v
	return s
}

// SetState sets the State field's value.
func (s *RecordForlistApplicationOutput) SetState(v string) *RecordForlistApplicationOutput {
	s.State = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *RecordForlistApplicationOutput) SetUserId(v string) *RecordForlistApplicationOutput {
	s.UserId = &v
	return s
}

// SetVersionName sets the VersionName field's value.
func (s *RecordForlistApplicationOutput) SetVersionName(v string) *RecordForlistApplicationOutput {
	s.VersionName = &v
	return s
}

const (
	// EnumOfStateForlistApplicationInputDeployed is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputDeployed = "DEPLOYED"

	// EnumOfStateForlistApplicationInputCreated is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputCreated = "CREATED"

	// EnumOfStateForlistApplicationInputStarting is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputStarting = "STARTING"

	// EnumOfStateForlistApplicationInputRunning is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputRunning = "RUNNING"

	// EnumOfStateForlistApplicationInputFailed is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputFailed = "FAILED"

	// EnumOfStateForlistApplicationInputCancelling is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputCancelling = "CANCELLING"

	// EnumOfStateForlistApplicationInputSucceeded is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputSucceeded = "SUCCEEDED"

	// EnumOfStateForlistApplicationInputStopped is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputStopped = "STOPPED"

	// EnumOfStateForlistApplicationInputUnknown is a EnumOfStateForlistApplicationInput enum value
	EnumOfStateForlistApplicationInputUnknown = "UNKNOWN"
)
