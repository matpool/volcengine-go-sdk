// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ecsiface provides an interface to enable mocking the ECS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecs

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
)

// ECSAPI provides an interface to enable mocking the
// ecs.ECS service client's API operation,
//
//	// volcengine sdk func uses an SDK service client to make a request to
//	// ECS.
//	func myFunc(svc ECSAPI) bool {
//	    // Make svc.AssociateInstancesIamRole request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := ecs.New(sess)
//
//	    myFunc(svc)
//	}
type ECSAPI interface {
	AssociateInstancesIamRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateInstancesIamRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateInstancesIamRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateInstancesIamRole(*AssociateInstancesIamRoleInput) (*AssociateInstancesIamRoleOutput, error)
	AssociateInstancesIamRoleWithContext(volcengine.Context, *AssociateInstancesIamRoleInput, ...request.Option) (*AssociateInstancesIamRoleOutput, error)
	AssociateInstancesIamRoleRequest(*AssociateInstancesIamRoleInput) (*request.Request, *AssociateInstancesIamRoleOutput)

	AttachKeyPairCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachKeyPairCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachKeyPairCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachKeyPair(*AttachKeyPairInput) (*AttachKeyPairOutput, error)
	AttachKeyPairWithContext(volcengine.Context, *AttachKeyPairInput, ...request.Option) (*AttachKeyPairOutput, error)
	AttachKeyPairRequest(*AttachKeyPairInput) (*request.Request, *AttachKeyPairOutput)

	CopyImageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CopyImageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CopyImageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CopyImage(*CopyImageInput) (*CopyImageOutput, error)
	CopyImageWithContext(volcengine.Context, *CopyImageInput, ...request.Option) (*CopyImageOutput, error)
	CopyImageRequest(*CopyImageInput) (*request.Request, *CopyImageOutput)

	CreateDeploymentSetCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDeploymentSetCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDeploymentSetCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDeploymentSet(*CreateDeploymentSetInput) (*CreateDeploymentSetOutput, error)
	CreateDeploymentSetWithContext(volcengine.Context, *CreateDeploymentSetInput, ...request.Option) (*CreateDeploymentSetOutput, error)
	CreateDeploymentSetRequest(*CreateDeploymentSetInput) (*request.Request, *CreateDeploymentSetOutput)

	CreateImageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateImageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateImageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateImage(*CreateImageInput) (*CreateImageOutput, error)
	CreateImageWithContext(volcengine.Context, *CreateImageInput, ...request.Option) (*CreateImageOutput, error)
	CreateImageRequest(*CreateImageInput) (*request.Request, *CreateImageOutput)

	CreateKeyPairCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyPairCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyPairCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateKeyPair(*CreateKeyPairInput) (*CreateKeyPairOutput, error)
	CreateKeyPairWithContext(volcengine.Context, *CreateKeyPairInput, ...request.Option) (*CreateKeyPairOutput, error)
	CreateKeyPairRequest(*CreateKeyPairInput) (*request.Request, *CreateKeyPairOutput)

	CreateSubscriptionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSubscriptionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSubscriptionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSubscription(*CreateSubscriptionInput) (*CreateSubscriptionOutput, error)
	CreateSubscriptionWithContext(volcengine.Context, *CreateSubscriptionInput, ...request.Option) (*CreateSubscriptionOutput, error)
	CreateSubscriptionRequest(*CreateSubscriptionInput) (*request.Request, *CreateSubscriptionOutput)

	CreateTagsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTagsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTagsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateTags(*CreateTagsInput) (*CreateTagsOutput, error)
	CreateTagsWithContext(volcengine.Context, *CreateTagsInput, ...request.Option) (*CreateTagsOutput, error)
	CreateTagsRequest(*CreateTagsInput) (*request.Request, *CreateTagsOutput)

	DeleteDeploymentSetCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDeploymentSetCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDeploymentSetCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDeploymentSet(*DeleteDeploymentSetInput) (*DeleteDeploymentSetOutput, error)
	DeleteDeploymentSetWithContext(volcengine.Context, *DeleteDeploymentSetInput, ...request.Option) (*DeleteDeploymentSetOutput, error)
	DeleteDeploymentSetRequest(*DeleteDeploymentSetInput) (*request.Request, *DeleteDeploymentSetOutput)

	DeleteImagesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteImagesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteImagesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteImages(*DeleteImagesInput) (*DeleteImagesOutput, error)
	DeleteImagesWithContext(volcengine.Context, *DeleteImagesInput, ...request.Option) (*DeleteImagesOutput, error)
	DeleteImagesRequest(*DeleteImagesInput) (*request.Request, *DeleteImagesOutput)

	DeleteInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInstance(*DeleteInstanceInput) (*DeleteInstanceOutput, error)
	DeleteInstanceWithContext(volcengine.Context, *DeleteInstanceInput, ...request.Option) (*DeleteInstanceOutput, error)
	DeleteInstanceRequest(*DeleteInstanceInput) (*request.Request, *DeleteInstanceOutput)

	DeleteInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInstances(*DeleteInstancesInput) (*DeleteInstancesOutput, error)
	DeleteInstancesWithContext(volcengine.Context, *DeleteInstancesInput, ...request.Option) (*DeleteInstancesOutput, error)
	DeleteInstancesRequest(*DeleteInstancesInput) (*request.Request, *DeleteInstancesOutput)

	DeleteKeyPairsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyPairsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyPairsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKeyPairs(*DeleteKeyPairsInput) (*DeleteKeyPairsOutput, error)
	DeleteKeyPairsWithContext(volcengine.Context, *DeleteKeyPairsInput, ...request.Option) (*DeleteKeyPairsOutput, error)
	DeleteKeyPairsRequest(*DeleteKeyPairsInput) (*request.Request, *DeleteKeyPairsOutput)

	DeleteTagsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteTagsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteTagsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteTags(*DeleteTagsInput) (*DeleteTagsOutput, error)
	DeleteTagsWithContext(volcengine.Context, *DeleteTagsInput, ...request.Option) (*DeleteTagsOutput, error)
	DeleteTagsRequest(*DeleteTagsInput) (*request.Request, *DeleteTagsOutput)

	DescribeAvailableResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailableResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailableResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailableResource(*DescribeAvailableResourceInput) (*DescribeAvailableResourceOutput, error)
	DescribeAvailableResourceWithContext(volcengine.Context, *DescribeAvailableResourceInput, ...request.Option) (*DescribeAvailableResourceOutput, error)
	DescribeAvailableResourceRequest(*DescribeAvailableResourceInput) (*request.Request, *DescribeAvailableResourceOutput)

	DescribeDeploymentSetSupportedInstanceTypeFamilyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDeploymentSetSupportedInstanceTypeFamilyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDeploymentSetSupportedInstanceTypeFamilyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDeploymentSetSupportedInstanceTypeFamily(*DescribeDeploymentSetSupportedInstanceTypeFamilyInput) (*DescribeDeploymentSetSupportedInstanceTypeFamilyOutput, error)
	DescribeDeploymentSetSupportedInstanceTypeFamilyWithContext(volcengine.Context, *DescribeDeploymentSetSupportedInstanceTypeFamilyInput, ...request.Option) (*DescribeDeploymentSetSupportedInstanceTypeFamilyOutput, error)
	DescribeDeploymentSetSupportedInstanceTypeFamilyRequest(*DescribeDeploymentSetSupportedInstanceTypeFamilyInput) (*request.Request, *DescribeDeploymentSetSupportedInstanceTypeFamilyOutput)

	DescribeDeploymentSetsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDeploymentSetsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDeploymentSetsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDeploymentSets(*DescribeDeploymentSetsInput) (*DescribeDeploymentSetsOutput, error)
	DescribeDeploymentSetsWithContext(volcengine.Context, *DescribeDeploymentSetsInput, ...request.Option) (*DescribeDeploymentSetsOutput, error)
	DescribeDeploymentSetsRequest(*DescribeDeploymentSetsInput) (*request.Request, *DescribeDeploymentSetsOutput)

	DescribeEventTypesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEventTypesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEventTypesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEventTypes(*DescribeEventTypesInput) (*DescribeEventTypesOutput, error)
	DescribeEventTypesWithContext(volcengine.Context, *DescribeEventTypesInput, ...request.Option) (*DescribeEventTypesOutput, error)
	DescribeEventTypesRequest(*DescribeEventTypesInput) (*request.Request, *DescribeEventTypesOutput)

	DescribeImageSharePermissionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImageSharePermissionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImageSharePermissionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImageSharePermission(*DescribeImageSharePermissionInput) (*DescribeImageSharePermissionOutput, error)
	DescribeImageSharePermissionWithContext(volcengine.Context, *DescribeImageSharePermissionInput, ...request.Option) (*DescribeImageSharePermissionOutput, error)
	DescribeImageSharePermissionRequest(*DescribeImageSharePermissionInput) (*request.Request, *DescribeImageSharePermissionOutput)

	DescribeImagesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImagesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImagesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImages(*DescribeImagesInput) (*DescribeImagesOutput, error)
	DescribeImagesWithContext(volcengine.Context, *DescribeImagesInput, ...request.Option) (*DescribeImagesOutput, error)
	DescribeImagesRequest(*DescribeImagesInput) (*request.Request, *DescribeImagesOutput)

	DescribeInstanceECSTerminalUrlCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceECSTerminalUrlCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceECSTerminalUrlCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceECSTerminalUrl(*DescribeInstanceECSTerminalUrlInput) (*DescribeInstanceECSTerminalUrlOutput, error)
	DescribeInstanceECSTerminalUrlWithContext(volcengine.Context, *DescribeInstanceECSTerminalUrlInput, ...request.Option) (*DescribeInstanceECSTerminalUrlOutput, error)
	DescribeInstanceECSTerminalUrlRequest(*DescribeInstanceECSTerminalUrlInput) (*request.Request, *DescribeInstanceECSTerminalUrlOutput)

	DescribeInstanceTypeFamiliesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceTypeFamiliesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceTypeFamiliesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceTypeFamilies(*DescribeInstanceTypeFamiliesInput) (*DescribeInstanceTypeFamiliesOutput, error)
	DescribeInstanceTypeFamiliesWithContext(volcengine.Context, *DescribeInstanceTypeFamiliesInput, ...request.Option) (*DescribeInstanceTypeFamiliesOutput, error)
	DescribeInstanceTypeFamiliesRequest(*DescribeInstanceTypeFamiliesInput) (*request.Request, *DescribeInstanceTypeFamiliesOutput)

	DescribeInstanceTypesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceTypesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceTypesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceTypes(*DescribeInstanceTypesInput) (*DescribeInstanceTypesOutput, error)
	DescribeInstanceTypesWithContext(volcengine.Context, *DescribeInstanceTypesInput, ...request.Option) (*DescribeInstanceTypesOutput, error)
	DescribeInstanceTypesRequest(*DescribeInstanceTypesInput) (*request.Request, *DescribeInstanceTypesOutput)

	DescribeInstanceVncUrlCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceVncUrlCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceVncUrlCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceVncUrl(*DescribeInstanceVncUrlInput) (*DescribeInstanceVncUrlOutput, error)
	DescribeInstanceVncUrlWithContext(volcengine.Context, *DescribeInstanceVncUrlInput, ...request.Option) (*DescribeInstanceVncUrlOutput, error)
	DescribeInstanceVncUrlRequest(*DescribeInstanceVncUrlInput) (*request.Request, *DescribeInstanceVncUrlOutput)

	DescribeInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstances(*DescribeInstancesInput) (*DescribeInstancesOutput, error)
	DescribeInstancesWithContext(volcengine.Context, *DescribeInstancesInput, ...request.Option) (*DescribeInstancesOutput, error)
	DescribeInstancesRequest(*DescribeInstancesInput) (*request.Request, *DescribeInstancesOutput)

	DescribeInstancesIamRolesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancesIamRolesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancesIamRolesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstancesIamRoles(*DescribeInstancesIamRolesInput) (*DescribeInstancesIamRolesOutput, error)
	DescribeInstancesIamRolesWithContext(volcengine.Context, *DescribeInstancesIamRolesInput, ...request.Option) (*DescribeInstancesIamRolesOutput, error)
	DescribeInstancesIamRolesRequest(*DescribeInstancesIamRolesInput) (*request.Request, *DescribeInstancesIamRolesOutput)

	DescribeKeyPairsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeyPairsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeyPairsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeyPairs(*DescribeKeyPairsInput) (*DescribeKeyPairsOutput, error)
	DescribeKeyPairsWithContext(volcengine.Context, *DescribeKeyPairsInput, ...request.Option) (*DescribeKeyPairsOutput, error)
	DescribeKeyPairsRequest(*DescribeKeyPairsInput) (*request.Request, *DescribeKeyPairsOutput)

	DescribeSpotAdviceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSpotAdviceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSpotAdviceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSpotAdvice(*DescribeSpotAdviceInput) (*DescribeSpotAdviceOutput, error)
	DescribeSpotAdviceWithContext(volcengine.Context, *DescribeSpotAdviceInput, ...request.Option) (*DescribeSpotAdviceOutput, error)
	DescribeSpotAdviceRequest(*DescribeSpotAdviceInput) (*request.Request, *DescribeSpotAdviceOutput)

	DescribeSubscriptionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSubscriptionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSubscriptionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSubscriptions(*DescribeSubscriptionsInput) (*DescribeSubscriptionsOutput, error)
	DescribeSubscriptionsWithContext(volcengine.Context, *DescribeSubscriptionsInput, ...request.Option) (*DescribeSubscriptionsOutput, error)
	DescribeSubscriptionsRequest(*DescribeSubscriptionsInput) (*request.Request, *DescribeSubscriptionsOutput)

	DescribeSystemEventsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSystemEventsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSystemEventsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSystemEvents(*DescribeSystemEventsInput) (*DescribeSystemEventsOutput, error)
	DescribeSystemEventsWithContext(volcengine.Context, *DescribeSystemEventsInput, ...request.Option) (*DescribeSystemEventsOutput, error)
	DescribeSystemEventsRequest(*DescribeSystemEventsInput) (*request.Request, *DescribeSystemEventsOutput)

	DescribeTagsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTagsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTagsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTags(*DescribeTagsInput) (*DescribeTagsOutput, error)
	DescribeTagsWithContext(volcengine.Context, *DescribeTagsInput, ...request.Option) (*DescribeTagsOutput, error)
	DescribeTagsRequest(*DescribeTagsInput) (*request.Request, *DescribeTagsOutput)

	DescribeTasksCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTasksCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTasksCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTasks(*DescribeTasksInput) (*DescribeTasksOutput, error)
	DescribeTasksWithContext(volcengine.Context, *DescribeTasksInput, ...request.Option) (*DescribeTasksOutput, error)
	DescribeTasksRequest(*DescribeTasksInput) (*request.Request, *DescribeTasksOutput)

	DescribeUserDataCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeUserDataCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeUserDataCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeUserData(*DescribeUserDataInput) (*DescribeUserDataOutput, error)
	DescribeUserDataWithContext(volcengine.Context, *DescribeUserDataInput, ...request.Option) (*DescribeUserDataOutput, error)
	DescribeUserDataRequest(*DescribeUserDataInput) (*request.Request, *DescribeUserDataOutput)

	DescribeZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeZones(*DescribeZonesInput) (*DescribeZonesOutput, error)
	DescribeZonesWithContext(volcengine.Context, *DescribeZonesInput, ...request.Option) (*DescribeZonesOutput, error)
	DescribeZonesRequest(*DescribeZonesInput) (*request.Request, *DescribeZonesOutput)

	DetachKeyPairCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachKeyPairCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachKeyPairCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachKeyPair(*DetachKeyPairInput) (*DetachKeyPairOutput, error)
	DetachKeyPairWithContext(volcengine.Context, *DetachKeyPairInput, ...request.Option) (*DetachKeyPairOutput, error)
	DetachKeyPairRequest(*DetachKeyPairInput) (*request.Request, *DetachKeyPairOutput)

	DisassociateInstancesIamRoleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateInstancesIamRoleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateInstancesIamRoleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateInstancesIamRole(*DisassociateInstancesIamRoleInput) (*DisassociateInstancesIamRoleOutput, error)
	DisassociateInstancesIamRoleWithContext(volcengine.Context, *DisassociateInstancesIamRoleInput, ...request.Option) (*DisassociateInstancesIamRoleOutput, error)
	DisassociateInstancesIamRoleRequest(*DisassociateInstancesIamRoleInput) (*request.Request, *DisassociateInstancesIamRoleOutput)

	ExportImageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ExportImageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ExportImageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ExportImage(*ExportImageInput) (*ExportImageOutput, error)
	ExportImageWithContext(volcengine.Context, *ExportImageInput, ...request.Option) (*ExportImageOutput, error)
	ExportImageRequest(*ExportImageInput) (*request.Request, *ExportImageOutput)

	GetConsoleOutputCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetConsoleOutputCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetConsoleOutputCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetConsoleOutput(*GetConsoleOutputInput) (*GetConsoleOutputOutput, error)
	GetConsoleOutputWithContext(volcengine.Context, *GetConsoleOutputInput, ...request.Option) (*GetConsoleOutputOutput, error)
	GetConsoleOutputRequest(*GetConsoleOutputInput) (*request.Request, *GetConsoleOutputOutput)

	GetConsoleScreenshotCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetConsoleScreenshotCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetConsoleScreenshotCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetConsoleScreenshot(*GetConsoleScreenshotInput) (*GetConsoleScreenshotOutput, error)
	GetConsoleScreenshotWithContext(volcengine.Context, *GetConsoleScreenshotInput, ...request.Option) (*GetConsoleScreenshotOutput, error)
	GetConsoleScreenshotRequest(*GetConsoleScreenshotInput) (*request.Request, *GetConsoleScreenshotOutput)

	ImportImageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ImportImageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportImageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportImage(*ImportImageInput) (*ImportImageOutput, error)
	ImportImageWithContext(volcengine.Context, *ImportImageInput, ...request.Option) (*ImportImageOutput, error)
	ImportImageRequest(*ImportImageInput) (*request.Request, *ImportImageOutput)

	ImportKeyPairCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ImportKeyPairCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportKeyPairCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportKeyPair(*ImportKeyPairInput) (*ImportKeyPairOutput, error)
	ImportKeyPairWithContext(volcengine.Context, *ImportKeyPairInput, ...request.Option) (*ImportKeyPairOutput, error)
	ImportKeyPairRequest(*ImportKeyPairInput) (*request.Request, *ImportKeyPairOutput)

	ModifyDeploymentSetAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDeploymentSetAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDeploymentSetAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDeploymentSetAttribute(*ModifyDeploymentSetAttributeInput) (*ModifyDeploymentSetAttributeOutput, error)
	ModifyDeploymentSetAttributeWithContext(volcengine.Context, *ModifyDeploymentSetAttributeInput, ...request.Option) (*ModifyDeploymentSetAttributeOutput, error)
	ModifyDeploymentSetAttributeRequest(*ModifyDeploymentSetAttributeInput) (*request.Request, *ModifyDeploymentSetAttributeOutput)

	ModifyImageAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyImageAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyImageAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyImageAttribute(*ModifyImageAttributeInput) (*ModifyImageAttributeOutput, error)
	ModifyImageAttributeWithContext(volcengine.Context, *ModifyImageAttributeInput, ...request.Option) (*ModifyImageAttributeOutput, error)
	ModifyImageAttributeRequest(*ModifyImageAttributeInput) (*request.Request, *ModifyImageAttributeOutput)

	ModifyImageSharePermissionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyImageSharePermissionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyImageSharePermissionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyImageSharePermission(*ModifyImageSharePermissionInput) (*ModifyImageSharePermissionOutput, error)
	ModifyImageSharePermissionWithContext(volcengine.Context, *ModifyImageSharePermissionInput, ...request.Option) (*ModifyImageSharePermissionOutput, error)
	ModifyImageSharePermissionRequest(*ModifyImageSharePermissionInput) (*request.Request, *ModifyImageSharePermissionOutput)

	ModifyInstanceAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceAttribute(*ModifyInstanceAttributeInput) (*ModifyInstanceAttributeOutput, error)
	ModifyInstanceAttributeWithContext(volcengine.Context, *ModifyInstanceAttributeInput, ...request.Option) (*ModifyInstanceAttributeOutput, error)
	ModifyInstanceAttributeRequest(*ModifyInstanceAttributeInput) (*request.Request, *ModifyInstanceAttributeOutput)

	ModifyInstanceChargeTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceChargeTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceChargeTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceChargeType(*ModifyInstanceChargeTypeInput) (*ModifyInstanceChargeTypeOutput, error)
	ModifyInstanceChargeTypeWithContext(volcengine.Context, *ModifyInstanceChargeTypeInput, ...request.Option) (*ModifyInstanceChargeTypeOutput, error)
	ModifyInstanceChargeTypeRequest(*ModifyInstanceChargeTypeInput) (*request.Request, *ModifyInstanceChargeTypeOutput)

	ModifyInstanceDeploymentCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceDeploymentCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceDeploymentCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceDeployment(*ModifyInstanceDeploymentInput) (*ModifyInstanceDeploymentOutput, error)
	ModifyInstanceDeploymentWithContext(volcengine.Context, *ModifyInstanceDeploymentInput, ...request.Option) (*ModifyInstanceDeploymentOutput, error)
	ModifyInstanceDeploymentRequest(*ModifyInstanceDeploymentInput) (*request.Request, *ModifyInstanceDeploymentOutput)

	ModifyInstanceSpecCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceSpecCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceSpecCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceSpec(*ModifyInstanceSpecInput) (*ModifyInstanceSpecOutput, error)
	ModifyInstanceSpecWithContext(volcengine.Context, *ModifyInstanceSpecInput, ...request.Option) (*ModifyInstanceSpecOutput, error)
	ModifyInstanceSpecRequest(*ModifyInstanceSpecInput) (*request.Request, *ModifyInstanceSpecOutput)

	ModifyKeyPairAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyKeyPairAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyKeyPairAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyKeyPairAttribute(*ModifyKeyPairAttributeInput) (*ModifyKeyPairAttributeOutput, error)
	ModifyKeyPairAttributeWithContext(volcengine.Context, *ModifyKeyPairAttributeInput, ...request.Option) (*ModifyKeyPairAttributeOutput, error)
	ModifyKeyPairAttributeRequest(*ModifyKeyPairAttributeInput) (*request.Request, *ModifyKeyPairAttributeOutput)

	ModifySubscriptionEventTypesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySubscriptionEventTypesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySubscriptionEventTypesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySubscriptionEventTypes(*ModifySubscriptionEventTypesInput) (*ModifySubscriptionEventTypesOutput, error)
	ModifySubscriptionEventTypesWithContext(volcengine.Context, *ModifySubscriptionEventTypesInput, ...request.Option) (*ModifySubscriptionEventTypesOutput, error)
	ModifySubscriptionEventTypesRequest(*ModifySubscriptionEventTypesInput) (*request.Request, *ModifySubscriptionEventTypesOutput)

	RebootInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RebootInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RebootInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RebootInstance(*RebootInstanceInput) (*RebootInstanceOutput, error)
	RebootInstanceWithContext(volcengine.Context, *RebootInstanceInput, ...request.Option) (*RebootInstanceOutput, error)
	RebootInstanceRequest(*RebootInstanceInput) (*request.Request, *RebootInstanceOutput)

	RebootInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RebootInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RebootInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RebootInstances(*RebootInstancesInput) (*RebootInstancesOutput, error)
	RebootInstancesWithContext(volcengine.Context, *RebootInstancesInput, ...request.Option) (*RebootInstancesOutput, error)
	RebootInstancesRequest(*RebootInstancesInput) (*request.Request, *RebootInstancesOutput)

	RenewInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RenewInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenewInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenewInstance(*RenewInstanceInput) (*RenewInstanceOutput, error)
	RenewInstanceWithContext(volcengine.Context, *RenewInstanceInput, ...request.Option) (*RenewInstanceOutput, error)
	RenewInstanceRequest(*RenewInstanceInput) (*request.Request, *RenewInstanceOutput)

	ReplaceSystemVolumeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReplaceSystemVolumeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReplaceSystemVolumeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReplaceSystemVolume(*ReplaceSystemVolumeInput) (*ReplaceSystemVolumeOutput, error)
	ReplaceSystemVolumeWithContext(volcengine.Context, *ReplaceSystemVolumeInput, ...request.Option) (*ReplaceSystemVolumeOutput, error)
	ReplaceSystemVolumeRequest(*ReplaceSystemVolumeInput) (*request.Request, *ReplaceSystemVolumeOutput)

	RunInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RunInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RunInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RunInstances(*RunInstancesInput) (*RunInstancesOutput, error)
	RunInstancesWithContext(volcengine.Context, *RunInstancesInput, ...request.Option) (*RunInstancesOutput, error)
	RunInstancesRequest(*RunInstancesInput) (*request.Request, *RunInstancesOutput)

	StartInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StartInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartInstance(*StartInstanceInput) (*StartInstanceOutput, error)
	StartInstanceWithContext(volcengine.Context, *StartInstanceInput, ...request.Option) (*StartInstanceOutput, error)
	StartInstanceRequest(*StartInstanceInput) (*request.Request, *StartInstanceOutput)

	StartInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StartInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartInstances(*StartInstancesInput) (*StartInstancesOutput, error)
	StartInstancesWithContext(volcengine.Context, *StartInstancesInput, ...request.Option) (*StartInstancesOutput, error)
	StartInstancesRequest(*StartInstancesInput) (*request.Request, *StartInstancesOutput)

	StopInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopInstance(*StopInstanceInput) (*StopInstanceOutput, error)
	StopInstanceWithContext(volcengine.Context, *StopInstanceInput, ...request.Option) (*StopInstanceOutput, error)
	StopInstanceRequest(*StopInstanceInput) (*request.Request, *StopInstanceOutput)

	StopInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopInstances(*StopInstancesInput) (*StopInstancesOutput, error)
	StopInstancesWithContext(volcengine.Context, *StopInstancesInput, ...request.Option) (*StopInstancesOutput, error)
	StopInstancesRequest(*StopInstancesInput) (*request.Request, *StopInstancesOutput)

	UpdateSystemEventsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSystemEventsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSystemEventsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSystemEvents(*UpdateSystemEventsInput) (*UpdateSystemEventsOutput, error)
	UpdateSystemEventsWithContext(volcengine.Context, *UpdateSystemEventsInput, ...request.Option) (*UpdateSystemEventsOutput, error)
	UpdateSystemEventsRequest(*UpdateSystemEventsInput) (*request.Request, *UpdateSystemEventsOutput)
}

var _ ECSAPI = (*ECS)(nil)
