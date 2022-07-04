// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package autoscalingiface provides an interface to enable mocking the AUTO_SCALING service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// AUTOSCALINGAPI provides an interface to enable mocking the
// autoscaling.AUTOSCALING service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // AUTO_SCALING.
//    func myFunc(svc AUTOSCALINGAPI) bool {
//        // Make svc.AttachDBInstances request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := autoscaling.New(sess)
//
//        myFunc(svc)
//    }
//
type AUTOSCALINGAPI interface {
	AttachDBInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachDBInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachDBInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachDBInstances(*AttachDBInstancesInput) (*AttachDBInstancesOutput, error)
	AttachDBInstancesWithContext(volcengine.Context, *AttachDBInstancesInput, ...request.Option) (*AttachDBInstancesOutput, error)
	AttachDBInstancesRequest(*AttachDBInstancesInput) (*request.Request, *AttachDBInstancesOutput)

	AttachInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachInstances(*AttachInstancesInput) (*AttachInstancesOutput, error)
	AttachInstancesWithContext(volcengine.Context, *AttachInstancesInput, ...request.Option) (*AttachInstancesOutput, error)
	AttachInstancesRequest(*AttachInstancesInput) (*request.Request, *AttachInstancesOutput)

	AttachServerGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachServerGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachServerGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachServerGroups(*AttachServerGroupsInput) (*AttachServerGroupsOutput, error)
	AttachServerGroupsWithContext(volcengine.Context, *AttachServerGroupsInput, ...request.Option) (*AttachServerGroupsOutput, error)
	AttachServerGroupsRequest(*AttachServerGroupsInput) (*request.Request, *AttachServerGroupsOutput)

	CompleteLifecycleActivityCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CompleteLifecycleActivityCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CompleteLifecycleActivityCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CompleteLifecycleActivity(*CompleteLifecycleActivityInput) (*CompleteLifecycleActivityOutput, error)
	CompleteLifecycleActivityWithContext(volcengine.Context, *CompleteLifecycleActivityInput, ...request.Option) (*CompleteLifecycleActivityOutput, error)
	CompleteLifecycleActivityRequest(*CompleteLifecycleActivityInput) (*request.Request, *CompleteLifecycleActivityOutput)

	CreateLifecycleHookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLifecycleHookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLifecycleHookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLifecycleHook(*CreateLifecycleHookInput) (*CreateLifecycleHookOutput, error)
	CreateLifecycleHookWithContext(volcengine.Context, *CreateLifecycleHookInput, ...request.Option) (*CreateLifecycleHookOutput, error)
	CreateLifecycleHookRequest(*CreateLifecycleHookInput) (*request.Request, *CreateLifecycleHookOutput)

	CreateScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateScalingConfigurationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateScalingConfiguration(*CreateScalingConfigurationInput) (*CreateScalingConfigurationOutput, error)
	CreateScalingConfigurationWithContext(volcengine.Context, *CreateScalingConfigurationInput, ...request.Option) (*CreateScalingConfigurationOutput, error)
	CreateScalingConfigurationRequest(*CreateScalingConfigurationInput) (*request.Request, *CreateScalingConfigurationOutput)

	CreateScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateScalingGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateScalingGroup(*CreateScalingGroupInput) (*CreateScalingGroupOutput, error)
	CreateScalingGroupWithContext(volcengine.Context, *CreateScalingGroupInput, ...request.Option) (*CreateScalingGroupOutput, error)
	CreateScalingGroupRequest(*CreateScalingGroupInput) (*request.Request, *CreateScalingGroupOutput)

	CreateScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateScalingPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateScalingPolicy(*CreateScalingPolicyInput) (*CreateScalingPolicyOutput, error)
	CreateScalingPolicyWithContext(volcengine.Context, *CreateScalingPolicyInput, ...request.Option) (*CreateScalingPolicyOutput, error)
	CreateScalingPolicyRequest(*CreateScalingPolicyInput) (*request.Request, *CreateScalingPolicyOutput)

	DeleteLifecycleHookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLifecycleHookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLifecycleHookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLifecycleHook(*DeleteLifecycleHookInput) (*DeleteLifecycleHookOutput, error)
	DeleteLifecycleHookWithContext(volcengine.Context, *DeleteLifecycleHookInput, ...request.Option) (*DeleteLifecycleHookOutput, error)
	DeleteLifecycleHookRequest(*DeleteLifecycleHookInput) (*request.Request, *DeleteLifecycleHookOutput)

	DeleteScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteScalingConfigurationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteScalingConfiguration(*DeleteScalingConfigurationInput) (*DeleteScalingConfigurationOutput, error)
	DeleteScalingConfigurationWithContext(volcengine.Context, *DeleteScalingConfigurationInput, ...request.Option) (*DeleteScalingConfigurationOutput, error)
	DeleteScalingConfigurationRequest(*DeleteScalingConfigurationInput) (*request.Request, *DeleteScalingConfigurationOutput)

	DeleteScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteScalingGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteScalingGroup(*DeleteScalingGroupInput) (*DeleteScalingGroupOutput, error)
	DeleteScalingGroupWithContext(volcengine.Context, *DeleteScalingGroupInput, ...request.Option) (*DeleteScalingGroupOutput, error)
	DeleteScalingGroupRequest(*DeleteScalingGroupInput) (*request.Request, *DeleteScalingGroupOutput)

	DeleteScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteScalingPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteScalingPolicy(*DeleteScalingPolicyInput) (*DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyWithContext(volcengine.Context, *DeleteScalingPolicyInput, ...request.Option) (*DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyRequest(*DeleteScalingPolicyInput) (*request.Request, *DeleteScalingPolicyOutput)

	DescribeLifecycleActivitiesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLifecycleActivitiesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLifecycleActivitiesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLifecycleActivities(*DescribeLifecycleActivitiesInput) (*DescribeLifecycleActivitiesOutput, error)
	DescribeLifecycleActivitiesWithContext(volcengine.Context, *DescribeLifecycleActivitiesInput, ...request.Option) (*DescribeLifecycleActivitiesOutput, error)
	DescribeLifecycleActivitiesRequest(*DescribeLifecycleActivitiesInput) (*request.Request, *DescribeLifecycleActivitiesOutput)

	DescribeLifecycleHooksCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLifecycleHooksCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLifecycleHooksCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLifecycleHooks(*DescribeLifecycleHooksInput) (*DescribeLifecycleHooksOutput, error)
	DescribeLifecycleHooksWithContext(volcengine.Context, *DescribeLifecycleHooksInput, ...request.Option) (*DescribeLifecycleHooksOutput, error)
	DescribeLifecycleHooksRequest(*DescribeLifecycleHooksInput) (*request.Request, *DescribeLifecycleHooksOutput)

	DescribeScalingActivitiesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingActivitiesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingActivitiesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingActivities(*DescribeScalingActivitiesInput) (*DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesWithContext(volcengine.Context, *DescribeScalingActivitiesInput, ...request.Option) (*DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesRequest(*DescribeScalingActivitiesInput) (*request.Request, *DescribeScalingActivitiesOutput)

	DescribeScalingConfigurationsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingConfigurationsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingConfigurationsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingConfigurations(*DescribeScalingConfigurationsInput) (*DescribeScalingConfigurationsOutput, error)
	DescribeScalingConfigurationsWithContext(volcengine.Context, *DescribeScalingConfigurationsInput, ...request.Option) (*DescribeScalingConfigurationsOutput, error)
	DescribeScalingConfigurationsRequest(*DescribeScalingConfigurationsInput) (*request.Request, *DescribeScalingConfigurationsOutput)

	DescribeScalingGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingGroups(*DescribeScalingGroupsInput) (*DescribeScalingGroupsOutput, error)
	DescribeScalingGroupsWithContext(volcengine.Context, *DescribeScalingGroupsInput, ...request.Option) (*DescribeScalingGroupsOutput, error)
	DescribeScalingGroupsRequest(*DescribeScalingGroupsInput) (*request.Request, *DescribeScalingGroupsOutput)

	DescribeScalingInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingInstances(*DescribeScalingInstancesInput) (*DescribeScalingInstancesOutput, error)
	DescribeScalingInstancesWithContext(volcengine.Context, *DescribeScalingInstancesInput, ...request.Option) (*DescribeScalingInstancesOutput, error)
	DescribeScalingInstancesRequest(*DescribeScalingInstancesInput) (*request.Request, *DescribeScalingInstancesOutput)

	DescribeScalingPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingPolicies(*DescribeScalingPoliciesInput) (*DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesWithContext(volcengine.Context, *DescribeScalingPoliciesInput, ...request.Option) (*DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesRequest(*DescribeScalingPoliciesInput) (*request.Request, *DescribeScalingPoliciesOutput)

	DetachDBInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachDBInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachDBInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachDBInstances(*DetachDBInstancesInput) (*DetachDBInstancesOutput, error)
	DetachDBInstancesWithContext(volcengine.Context, *DetachDBInstancesInput, ...request.Option) (*DetachDBInstancesOutput, error)
	DetachDBInstancesRequest(*DetachDBInstancesInput) (*request.Request, *DetachDBInstancesOutput)

	DetachInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachInstances(*DetachInstancesInput) (*DetachInstancesOutput, error)
	DetachInstancesWithContext(volcengine.Context, *DetachInstancesInput, ...request.Option) (*DetachInstancesOutput, error)
	DetachInstancesRequest(*DetachInstancesInput) (*request.Request, *DetachInstancesOutput)

	DetachServerGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachServerGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachServerGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachServerGroups(*DetachServerGroupsInput) (*DetachServerGroupsOutput, error)
	DetachServerGroupsWithContext(volcengine.Context, *DetachServerGroupsInput, ...request.Option) (*DetachServerGroupsOutput, error)
	DetachServerGroupsRequest(*DetachServerGroupsInput) (*request.Request, *DetachServerGroupsOutput)

	DisableScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableScalingGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableScalingGroup(*DisableScalingGroupInput) (*DisableScalingGroupOutput, error)
	DisableScalingGroupWithContext(volcengine.Context, *DisableScalingGroupInput, ...request.Option) (*DisableScalingGroupOutput, error)
	DisableScalingGroupRequest(*DisableScalingGroupInput) (*request.Request, *DisableScalingGroupOutput)

	DisableScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableScalingPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableScalingPolicy(*DisableScalingPolicyInput) (*DisableScalingPolicyOutput, error)
	DisableScalingPolicyWithContext(volcengine.Context, *DisableScalingPolicyInput, ...request.Option) (*DisableScalingPolicyOutput, error)
	DisableScalingPolicyRequest(*DisableScalingPolicyInput) (*request.Request, *DisableScalingPolicyOutput)

	EnableScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableScalingConfigurationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableScalingConfiguration(*EnableScalingConfigurationInput) (*EnableScalingConfigurationOutput, error)
	EnableScalingConfigurationWithContext(volcengine.Context, *EnableScalingConfigurationInput, ...request.Option) (*EnableScalingConfigurationOutput, error)
	EnableScalingConfigurationRequest(*EnableScalingConfigurationInput) (*request.Request, *EnableScalingConfigurationOutput)

	EnableScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableScalingGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableScalingGroup(*EnableScalingGroupInput) (*EnableScalingGroupOutput, error)
	EnableScalingGroupWithContext(volcengine.Context, *EnableScalingGroupInput, ...request.Option) (*EnableScalingGroupOutput, error)
	EnableScalingGroupRequest(*EnableScalingGroupInput) (*request.Request, *EnableScalingGroupOutput)

	EnableScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableScalingPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableScalingPolicy(*EnableScalingPolicyInput) (*EnableScalingPolicyOutput, error)
	EnableScalingPolicyWithContext(volcengine.Context, *EnableScalingPolicyInput, ...request.Option) (*EnableScalingPolicyOutput, error)
	EnableScalingPolicyRequest(*EnableScalingPolicyInput) (*request.Request, *EnableScalingPolicyOutput)

	ModifyLifecycleHookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLifecycleHookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLifecycleHookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLifecycleHook(*ModifyLifecycleHookInput) (*ModifyLifecycleHookOutput, error)
	ModifyLifecycleHookWithContext(volcengine.Context, *ModifyLifecycleHookInput, ...request.Option) (*ModifyLifecycleHookOutput, error)
	ModifyLifecycleHookRequest(*ModifyLifecycleHookInput) (*request.Request, *ModifyLifecycleHookOutput)

	ModifyScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyScalingConfigurationCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyScalingConfiguration(*ModifyScalingConfigurationInput) (*ModifyScalingConfigurationOutput, error)
	ModifyScalingConfigurationWithContext(volcengine.Context, *ModifyScalingConfigurationInput, ...request.Option) (*ModifyScalingConfigurationOutput, error)
	ModifyScalingConfigurationRequest(*ModifyScalingConfigurationInput) (*request.Request, *ModifyScalingConfigurationOutput)

	ModifyScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyScalingGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyScalingGroup(*ModifyScalingGroupInput) (*ModifyScalingGroupOutput, error)
	ModifyScalingGroupWithContext(volcengine.Context, *ModifyScalingGroupInput, ...request.Option) (*ModifyScalingGroupOutput, error)
	ModifyScalingGroupRequest(*ModifyScalingGroupInput) (*request.Request, *ModifyScalingGroupOutput)

	ModifyScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyScalingPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyScalingPolicy(*ModifyScalingPolicyInput) (*ModifyScalingPolicyOutput, error)
	ModifyScalingPolicyWithContext(volcengine.Context, *ModifyScalingPolicyInput, ...request.Option) (*ModifyScalingPolicyOutput, error)
	ModifyScalingPolicyRequest(*ModifyScalingPolicyInput) (*request.Request, *ModifyScalingPolicyOutput)

	RemoveInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveInstances(*RemoveInstancesInput) (*RemoveInstancesOutput, error)
	RemoveInstancesWithContext(volcengine.Context, *RemoveInstancesInput, ...request.Option) (*RemoveInstancesOutput, error)
	RemoveInstancesRequest(*RemoveInstancesInput) (*request.Request, *RemoveInstancesOutput)

	SetInstancesProtectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetInstancesProtectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetInstancesProtectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetInstancesProtection(*SetInstancesProtectionInput) (*SetInstancesProtectionOutput, error)
	SetInstancesProtectionWithContext(volcengine.Context, *SetInstancesProtectionInput, ...request.Option) (*SetInstancesProtectionOutput, error)
	SetInstancesProtectionRequest(*SetInstancesProtectionInput) (*request.Request, *SetInstancesProtectionOutput)
}

var _ AUTOSCALINGAPI = (*AUTOSCALING)(nil)
