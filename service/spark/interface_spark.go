// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sparkiface provides an interface to enable mocking the SPARK service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
)

// SPARKAPI provides an interface to enable mocking the
// spark.SPARK service client's API operation,
//
//	// volcengine sdk func uses an SDK service client to make a request to
//	// SPARK.
//	func myFunc(svc SPARKAPI) bool {
//	    // Make svc.CreateApplication request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := spark.New(sess)
//
//	    myFunc(svc)
//	}
type SPARKAPI interface {
	CreateApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateApplication(*CreateApplicationInput) (*CreateApplicationOutput, error)
	CreateApplicationWithContext(volcengine.Context, *CreateApplicationInput, ...request.Option) (*CreateApplicationOutput, error)
	CreateApplicationRequest(*CreateApplicationInput) (*request.Request, *CreateApplicationOutput)

	CreateProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateProject(*CreateProjectInput) (*CreateProjectOutput, error)
	CreateProjectWithContext(volcengine.Context, *CreateProjectInput, ...request.Option) (*CreateProjectOutput, error)
	CreateProjectRequest(*CreateProjectInput) (*request.Request, *CreateProjectOutput)

	CreateResourcePoolCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateResourcePoolCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateResourcePoolCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateResourcePool(*CreateResourcePoolInput) (*CreateResourcePoolOutput, error)
	CreateResourcePoolWithContext(volcengine.Context, *CreateResourcePoolInput, ...request.Option) (*CreateResourcePoolOutput, error)
	CreateResourcePoolRequest(*CreateResourcePoolInput) (*request.Request, *CreateResourcePoolOutput)

	DeleteApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteApplication(*DeleteApplicationInput) (*DeleteApplicationOutput, error)
	DeleteApplicationWithContext(volcengine.Context, *DeleteApplicationInput, ...request.Option) (*DeleteApplicationOutput, error)
	DeleteApplicationRequest(*DeleteApplicationInput) (*request.Request, *DeleteApplicationOutput)

	DeleteProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteProject(*DeleteProjectInput) (*DeleteProjectOutput, error)
	DeleteProjectWithContext(volcengine.Context, *DeleteProjectInput, ...request.Option) (*DeleteProjectOutput, error)
	DeleteProjectRequest(*DeleteProjectInput) (*request.Request, *DeleteProjectOutput)

	DeleteResourcePoolCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteResourcePoolCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteResourcePoolCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteResourcePool(*DeleteResourcePoolInput) (*DeleteResourcePoolOutput, error)
	DeleteResourcePoolWithContext(volcengine.Context, *DeleteResourcePoolInput, ...request.Option) (*DeleteResourcePoolOutput, error)
	DeleteResourcePoolRequest(*DeleteResourcePoolInput) (*request.Request, *DeleteResourcePoolOutput)

	DescribeApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeApplication(*DescribeApplicationInput) (*DescribeApplicationOutput, error)
	DescribeApplicationWithContext(volcengine.Context, *DescribeApplicationInput, ...request.Option) (*DescribeApplicationOutput, error)
	DescribeApplicationRequest(*DescribeApplicationInput) (*request.Request, *DescribeApplicationOutput)

	DescribeApplicationInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeApplicationInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeApplicationInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeApplicationInstance(*DescribeApplicationInstanceInput) (*DescribeApplicationInstanceOutput, error)
	DescribeApplicationInstanceWithContext(volcengine.Context, *DescribeApplicationInstanceInput, ...request.Option) (*DescribeApplicationInstanceOutput, error)
	DescribeApplicationInstanceRequest(*DescribeApplicationInstanceInput) (*request.Request, *DescribeApplicationInstanceOutput)

	DescribeProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeProject(*DescribeProjectInput) (*DescribeProjectOutput, error)
	DescribeProjectWithContext(volcengine.Context, *DescribeProjectInput, ...request.Option) (*DescribeProjectOutput, error)
	DescribeProjectRequest(*DescribeProjectInput) (*request.Request, *DescribeProjectOutput)

	DescribeResourcePoolCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeResourcePoolCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeResourcePoolCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeResourcePool(*DescribeResourcePoolInput) (*DescribeResourcePoolOutput, error)
	DescribeResourcePoolWithContext(volcengine.Context, *DescribeResourcePoolInput, ...request.Option) (*DescribeResourcePoolOutput, error)
	DescribeResourcePoolRequest(*DescribeResourcePoolInput) (*request.Request, *DescribeResourcePoolOutput)

	ExistResourcePoolCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ExistResourcePoolCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ExistResourcePoolCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ExistResourcePool(*ExistResourcePoolInput) (*ExistResourcePoolOutput, error)
	ExistResourcePoolWithContext(volcengine.Context, *ExistResourcePoolInput, ...request.Option) (*ExistResourcePoolOutput, error)
	ExistResourcePoolRequest(*ExistResourcePoolInput) (*request.Request, *ExistResourcePoolOutput)

	ListAppInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAppInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAppInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAppInstance(*ListAppInstanceInput) (*ListAppInstanceOutput, error)
	ListAppInstanceWithContext(volcengine.Context, *ListAppInstanceInput, ...request.Option) (*ListAppInstanceOutput, error)
	ListAppInstanceRequest(*ListAppInstanceInput) (*request.Request, *ListAppInstanceOutput)

	ListApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplication(*ListApplicationInput) (*ListApplicationOutput, error)
	ListApplicationWithContext(volcengine.Context, *ListApplicationInput, ...request.Option) (*ListApplicationOutput, error)
	ListApplicationRequest(*ListApplicationInput) (*request.Request, *ListApplicationOutput)

	ListApplicationHistoryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationHistoryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationHistoryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplicationHistory(*ListApplicationHistoryInput) (*ListApplicationHistoryOutput, error)
	ListApplicationHistoryWithContext(volcengine.Context, *ListApplicationHistoryInput, ...request.Option) (*ListApplicationHistoryOutput, error)
	ListApplicationHistoryRequest(*ListApplicationHistoryInput) (*request.Request, *ListApplicationHistoryOutput)

	ListProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListProject(*ListProjectInput) (*ListProjectOutput, error)
	ListProjectWithContext(volcengine.Context, *ListProjectInput, ...request.Option) (*ListProjectOutput, error)
	ListProjectRequest(*ListProjectInput) (*request.Request, *ListProjectOutput)

	ListResourcePoolCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListResourcePoolCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListResourcePoolCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListResourcePool(*ListResourcePoolInput) (*ListResourcePoolOutput, error)
	ListResourcePoolWithContext(volcengine.Context, *ListResourcePoolInput, ...request.Option) (*ListResourcePoolOutput, error)
	ListResourcePoolRequest(*ListResourcePoolInput) (*request.Request, *ListResourcePoolOutput)

	ListZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListZone(*ListZoneInput) (*ListZoneOutput, error)
	ListZoneWithContext(volcengine.Context, *ListZoneInput, ...request.Option) (*ListZoneOutput, error)
	ListZoneRequest(*ListZoneInput) (*request.Request, *ListZoneOutput)

	ModifyApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyApplication(*ModifyApplicationInput) (*ModifyApplicationOutput, error)
	ModifyApplicationWithContext(volcengine.Context, *ModifyApplicationInput, ...request.Option) (*ModifyApplicationOutput, error)
	ModifyApplicationRequest(*ModifyApplicationInput) (*request.Request, *ModifyApplicationOutput)

	StartApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StartApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartApplication(*StartApplicationInput) (*StartApplicationOutput, error)
	StartApplicationWithContext(volcengine.Context, *StartApplicationInput, ...request.Option) (*StartApplicationOutput, error)
	StartApplicationRequest(*StartApplicationInput) (*request.Request, *StartApplicationOutput)

	StopApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopApplication(*StopApplicationInput) (*StopApplicationOutput, error)
	StopApplicationWithContext(volcengine.Context, *StopApplicationInput, ...request.Option) (*StopApplicationOutput, error)
	StopApplicationRequest(*StopApplicationInput) (*request.Request, *StopApplicationOutput)

	UpdateProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateProject(*UpdateProjectInput) (*UpdateProjectOutput, error)
	UpdateProjectWithContext(volcengine.Context, *UpdateProjectInput, ...request.Option) (*UpdateProjectOutput, error)
	UpdateProjectRequest(*UpdateProjectInput) (*request.Request, *UpdateProjectOutput)
}

var _ SPARKAPI = (*SPARK)(nil)
