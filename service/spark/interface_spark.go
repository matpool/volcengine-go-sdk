// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sparkiface provides an interface to enable mocking the SPARK service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// SPARKAPI provides an interface to enable mocking the
// spark.SPARK service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // SPARK.
//    func myFunc(svc SPARKAPI) bool {
//        // Make svc.CreateApplication request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := spark.New(sess)
//
//        myFunc(svc)
//    }
//
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

	CreateResourcePoolOneStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateResourcePoolOneStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateResourcePoolOneStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateResourcePoolOneStep(*CreateResourcePoolOneStepInput) (*CreateResourcePoolOneStepOutput, error)
	CreateResourcePoolOneStepWithContext(volcengine.Context, *CreateResourcePoolOneStepInput, ...request.Option) (*CreateResourcePoolOneStepOutput, error)
	CreateResourcePoolOneStepRequest(*CreateResourcePoolOneStepInput) (*request.Request, *CreateResourcePoolOneStepOutput)

	DeleteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Delete(*DeleteInput) (*DeleteOutput, error)
	DeleteWithContext(volcengine.Context, *DeleteInput, ...request.Option) (*DeleteOutput, error)
	DeleteRequest(*DeleteInput) (*request.Request, *DeleteOutput)

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

	DetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	Detail(*DetailInput) (*DetailOutput, error)
	DetailWithContext(volcengine.Context, *DetailInput, ...request.Option) (*DetailOutput, error)
	DetailRequest(*DetailInput) (*request.Request, *DetailOutput)

	ExitResourcePoolCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ExitResourcePoolCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ExitResourcePoolCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ExitResourcePool(*ExitResourcePoolInput) (*ExitResourcePoolOutput, error)
	ExitResourcePoolWithContext(volcengine.Context, *ExitResourcePoolInput, ...request.Option) (*ExitResourcePoolOutput, error)
	ExitResourcePoolRequest(*ExitResourcePoolInput) (*request.Request, *ExitResourcePoolOutput)

	GetApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetApplication(*GetApplicationInput) (*GetApplicationOutput, error)
	GetApplicationWithContext(volcengine.Context, *GetApplicationInput, ...request.Option) (*GetApplicationOutput, error)
	GetApplicationRequest(*GetApplicationInput) (*request.Request, *GetApplicationOutput)

	ListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	List(*ListInput) (*ListOutput, error)
	ListWithContext(volcengine.Context, *ListInput, ...request.Option) (*ListOutput, error)
	ListRequest(*ListInput) (*request.Request, *ListOutput)

	ListApplicationHistoryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationHistoryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationHistoryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplicationHistory(*ListApplicationHistoryInput) (*ListApplicationHistoryOutput, error)
	ListApplicationHistoryWithContext(volcengine.Context, *ListApplicationHistoryInput, ...request.Option) (*ListApplicationHistoryOutput, error)
	ListApplicationHistoryRequest(*ListApplicationHistoryInput) (*request.Request, *ListApplicationHistoryOutput)

	ListApplicationInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListApplicationInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListApplicationInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListApplicationInstance(*ListApplicationInstanceInput) (*ListApplicationInstanceOutput, error)
	ListApplicationInstanceWithContext(volcengine.Context, *ListApplicationInstanceInput, ...request.Option) (*ListApplicationInstanceOutput, error)
	ListApplicationInstanceRequest(*ListApplicationInstanceInput) (*request.Request, *ListApplicationInstanceOutput)

	ListProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListProject(*ListProjectInput) (*ListProjectOutput, error)
	ListProjectWithContext(volcengine.Context, *ListProjectInput, ...request.Option) (*ListProjectOutput, error)
	ListProjectRequest(*ListProjectInput) (*request.Request, *ListProjectOutput)

	ListZoneCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListZoneCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListZoneCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListZone(*volcengineCommonQuery) (*volcengineCommonQuery, error)
	ListZoneWithContext(volcengine.Context, *volcengineCommonQuery, ...request.Option) (*volcengineCommonQuery, error)
	ListZoneRequest(*volcengineCommonQuery) (*request.Request, *volcengineCommonQuery)

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

	UpdateApplicationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateApplicationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateApplicationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateApplication(*UpdateApplicationInput) (*UpdateApplicationOutput, error)
	UpdateApplicationWithContext(volcengine.Context, *UpdateApplicationInput, ...request.Option) (*UpdateApplicationOutput, error)
	UpdateApplicationRequest(*UpdateApplicationInput) (*request.Request, *UpdateApplicationOutput)

	UpdateProjectCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateProjectCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateProjectCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateProject(*UpdateProjectInput) (*UpdateProjectOutput, error)
	UpdateProjectWithContext(volcengine.Context, *UpdateProjectInput, ...request.Option) (*UpdateProjectOutput, error)
	UpdateProjectRequest(*UpdateProjectInput) (*request.Request, *UpdateProjectOutput)
}

var _ SPARKAPI = (*SPARK)(nil)
