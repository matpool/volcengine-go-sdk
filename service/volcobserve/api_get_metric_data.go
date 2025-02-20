// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetMetricDataCommon = "GetMetricData"

// GetMetricDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetMetricDataCommon operation. The "output" return
// value will be populated with the GetMetricDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetMetricDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetMetricDataCommon Send returns without error.
//
// See GetMetricDataCommon for more information on using the GetMetricDataCommon
// API call, and error handling.
//
//	// Example sending a request using the GetMetricDataCommonRequest method.
//	req, resp := client.GetMetricDataCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VOLCOBSERVE) GetMetricDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetMetricDataCommon,
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

// GetMetricDataCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation GetMetricDataCommon for usage and error information.
func (c *VOLCOBSERVE) GetMetricDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetMetricDataCommonRequest(input)
	return out, req.Send()
}

// GetMetricDataCommonWithContext is the same as GetMetricDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetMetricDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) GetMetricDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetMetricDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetMetricData = "GetMetricData"

// GetMetricDataRequest generates a "volcengine/request.Request" representing the
// client's request for the GetMetricData operation. The "output" return
// value will be populated with the GetMetricDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetMetricDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetMetricDataCommon Send returns without error.
//
// See GetMetricData for more information on using the GetMetricData
// API call, and error handling.
//
//	// Example sending a request using the GetMetricDataRequest method.
//	req, resp := client.GetMetricDataRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VOLCOBSERVE) GetMetricDataRequest(input *GetMetricDataInput) (req *request.Request, output *GetMetricDataOutput) {
	op := &request.Operation{
		Name:       opGetMetricData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMetricDataInput{}
	}

	output = &GetMetricDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetMetricData API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation GetMetricData for usage and error information.
func (c *VOLCOBSERVE) GetMetricData(input *GetMetricDataInput) (*GetMetricDataOutput, error) {
	req, out := c.GetMetricDataRequest(input)
	return out, req.Send()
}

// GetMetricDataWithContext is the same as GetMetricData with the addition of
// the ability to pass a context and additional request options.
//
// See GetMetricData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) GetMetricDataWithContext(ctx volcengine.Context, input *GetMetricDataInput, opts ...request.Option) (*GetMetricDataOutput, error) {
	req, out := c.GetMetricDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetMetricDataOutput struct {
	_ struct{} `type:"structure"`

	DescriptionCN *string `type:"string"`

	DescriptionEN *string `type:"string"`

	EndTime *int64 `type:"integer"`

	MetricDataResults []*MetricDataResultForGetMetricDataOutput `type:"list"`

	MetricName *string `type:"string"`

	Namespace *string `type:"string"`

	Period *string `type:"string"`

	StartTime *int64 `type:"integer"`

	Unit *string `type:"string"`
}

// String returns the string representation
func (s DataForGetMetricDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetMetricDataOutput) GoString() string {
	return s.String()
}

// SetDescriptionCN sets the DescriptionCN field's value.
func (s *DataForGetMetricDataOutput) SetDescriptionCN(v string) *DataForGetMetricDataOutput {
	s.DescriptionCN = &v
	return s
}

// SetDescriptionEN sets the DescriptionEN field's value.
func (s *DataForGetMetricDataOutput) SetDescriptionEN(v string) *DataForGetMetricDataOutput {
	s.DescriptionEN = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DataForGetMetricDataOutput) SetEndTime(v int64) *DataForGetMetricDataOutput {
	s.EndTime = &v
	return s
}

// SetMetricDataResults sets the MetricDataResults field's value.
func (s *DataForGetMetricDataOutput) SetMetricDataResults(v []*MetricDataResultForGetMetricDataOutput) *DataForGetMetricDataOutput {
	s.MetricDataResults = v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *DataForGetMetricDataOutput) SetMetricName(v string) *DataForGetMetricDataOutput {
	s.MetricName = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DataForGetMetricDataOutput) SetNamespace(v string) *DataForGetMetricDataOutput {
	s.Namespace = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *DataForGetMetricDataOutput) SetPeriod(v string) *DataForGetMetricDataOutput {
	s.Period = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DataForGetMetricDataOutput) SetStartTime(v int64) *DataForGetMetricDataOutput {
	s.StartTime = &v
	return s
}

// SetUnit sets the Unit field's value.
func (s *DataForGetMetricDataOutput) SetUnit(v string) *DataForGetMetricDataOutput {
	s.Unit = &v
	return s
}

type DataPointForGetMetricDataOutput struct {
	_ struct{} `type:"structure"`

	Timestamp *int64 `type:"integer"`

	Value *float64 `type:"double"`
}

// String returns the string representation
func (s DataPointForGetMetricDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataPointForGetMetricDataOutput) GoString() string {
	return s.String()
}

// SetTimestamp sets the Timestamp field's value.
func (s *DataPointForGetMetricDataOutput) SetTimestamp(v int64) *DataPointForGetMetricDataOutput {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DataPointForGetMetricDataOutput) SetValue(v float64) *DataPointForGetMetricDataOutput {
	s.Value = &v
	return s
}

type DimensionForGetMetricDataInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s DimensionForGetMetricDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForGetMetricDataInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForGetMetricDataInput) SetName(v string) *DimensionForGetMetricDataInput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForGetMetricDataInput) SetValue(v string) *DimensionForGetMetricDataInput {
	s.Value = &v
	return s
}

type DimensionForGetMetricDataOutput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s DimensionForGetMetricDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForGetMetricDataOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForGetMetricDataOutput) SetName(v string) *DimensionForGetMetricDataOutput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForGetMetricDataOutput) SetValue(v string) *DimensionForGetMetricDataOutput {
	s.Value = &v
	return s
}

type GetMetricDataInput struct {
	_ struct{} `type:"structure"`

	EndTime *int64 `type:"integer"`

	Instances []*InstanceForGetMetricDataInput `type:"list"`

	MetricName *string `type:"string"`

	Namespace *string `type:"string"`

	Period *string `type:"string"`

	StartTime *int64 `type:"integer"`

	SubNamespace *string `type:"string"`
}

// String returns the string representation
func (s GetMetricDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMetricDataInput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *GetMetricDataInput) SetEndTime(v int64) *GetMetricDataInput {
	s.EndTime = &v
	return s
}

// SetInstances sets the Instances field's value.
func (s *GetMetricDataInput) SetInstances(v []*InstanceForGetMetricDataInput) *GetMetricDataInput {
	s.Instances = v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *GetMetricDataInput) SetMetricName(v string) *GetMetricDataInput {
	s.MetricName = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *GetMetricDataInput) SetNamespace(v string) *GetMetricDataInput {
	s.Namespace = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *GetMetricDataInput) SetPeriod(v string) *GetMetricDataInput {
	s.Period = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *GetMetricDataInput) SetStartTime(v int64) *GetMetricDataInput {
	s.StartTime = &v
	return s
}

// SetSubNamespace sets the SubNamespace field's value.
func (s *GetMetricDataInput) SetSubNamespace(v string) *GetMetricDataInput {
	s.SubNamespace = &v
	return s
}

type GetMetricDataOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data *DataForGetMetricDataOutput `type:"structure"`
}

// String returns the string representation
func (s GetMetricDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMetricDataOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetMetricDataOutput) SetData(v *DataForGetMetricDataOutput) *GetMetricDataOutput {
	s.Data = v
	return s
}

type InstanceForGetMetricDataInput struct {
	_ struct{} `type:"structure"`

	Dimensions []*DimensionForGetMetricDataInput `type:"list"`
}

// String returns the string representation
func (s InstanceForGetMetricDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForGetMetricDataInput) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *InstanceForGetMetricDataInput) SetDimensions(v []*DimensionForGetMetricDataInput) *InstanceForGetMetricDataInput {
	s.Dimensions = v
	return s
}

type MetricDataResultForGetMetricDataOutput struct {
	_ struct{} `type:"structure"`

	DataPoints []*DataPointForGetMetricDataOutput `type:"list"`

	Dimensions []*DimensionForGetMetricDataOutput `type:"list"`

	Legend *string `type:"string"`
}

// String returns the string representation
func (s MetricDataResultForGetMetricDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricDataResultForGetMetricDataOutput) GoString() string {
	return s.String()
}

// SetDataPoints sets the DataPoints field's value.
func (s *MetricDataResultForGetMetricDataOutput) SetDataPoints(v []*DataPointForGetMetricDataOutput) *MetricDataResultForGetMetricDataOutput {
	s.DataPoints = v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *MetricDataResultForGetMetricDataOutput) SetDimensions(v []*DimensionForGetMetricDataOutput) *MetricDataResultForGetMetricDataOutput {
	s.Dimensions = v
	return s
}

// SetLegend sets the Legend field's value.
func (s *MetricDataResultForGetMetricDataOutput) SetLegend(v string) *MetricDataResultForGetMetricDataOutput {
	s.Legend = &v
	return s
}
