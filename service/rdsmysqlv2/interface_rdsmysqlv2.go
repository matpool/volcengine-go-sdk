// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdsmysqlv2iface provides an interface to enable mocking the RDS_MYSQL_V2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// RDSMYSQLV2API provides an interface to enable mocking the
// rdsmysqlv2.RDSMYSQLV2 service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // RDS_MYSQL_V2.
//    func myFunc(svc RDSMYSQLV2API) bool {
//        // Make svc.AssociateAllowList request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rdsmysqlv2.New(sess)
//
//        myFunc(svc)
//    }
//
type RDSMYSQLV2API interface {
	AssociateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateAllowList(*AssociateAllowListInput) (*AssociateAllowListOutput, error)
	AssociateAllowListWithContext(volcengine.Context, *AssociateAllowListInput, ...request.Option) (*AssociateAllowListOutput, error)
	AssociateAllowListRequest(*AssociateAllowListInput) (*request.Request, *AssociateAllowListOutput)

	CreateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAllowList(*CreateAllowListInput) (*CreateAllowListOutput, error)
	CreateAllowListWithContext(volcengine.Context, *CreateAllowListInput, ...request.Option) (*CreateAllowListOutput, error)
	CreateAllowListRequest(*CreateAllowListInput) (*request.Request, *CreateAllowListOutput)

	CreateBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBackup(*CreateBackupInput) (*CreateBackupOutput, error)
	CreateBackupWithContext(volcengine.Context, *CreateBackupInput, ...request.Option) (*CreateBackupOutput, error)
	CreateBackupRequest(*CreateBackupInput) (*request.Request, *CreateBackupOutput)

	CreateDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBAccount(*CreateDBAccountInput) (*CreateDBAccountOutput, error)
	CreateDBAccountWithContext(volcengine.Context, *CreateDBAccountInput, ...request.Option) (*CreateDBAccountOutput, error)
	CreateDBAccountRequest(*CreateDBAccountInput) (*request.Request, *CreateDBAccountOutput)

	CreateDBEndpointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBEndpointCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBEndpointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBEndpoint(*CreateDBEndpointInput) (*CreateDBEndpointOutput, error)
	CreateDBEndpointWithContext(volcengine.Context, *CreateDBEndpointInput, ...request.Option) (*CreateDBEndpointOutput, error)
	CreateDBEndpointRequest(*CreateDBEndpointInput) (*request.Request, *CreateDBEndpointOutput)

	CreateDBEndpointPublicAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBEndpointPublicAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBEndpointPublicAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBEndpointPublicAddress(*CreateDBEndpointPublicAddressInput) (*CreateDBEndpointPublicAddressOutput, error)
	CreateDBEndpointPublicAddressWithContext(volcengine.Context, *CreateDBEndpointPublicAddressInput, ...request.Option) (*CreateDBEndpointPublicAddressOutput, error)
	CreateDBEndpointPublicAddressRequest(*CreateDBEndpointPublicAddressInput) (*request.Request, *CreateDBEndpointPublicAddressOutput)

	CreateDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBInstance(*CreateDBInstanceInput) (*CreateDBInstanceOutput, error)
	CreateDBInstanceWithContext(volcengine.Context, *CreateDBInstanceInput, ...request.Option) (*CreateDBInstanceOutput, error)
	CreateDBInstanceRequest(*CreateDBInstanceInput) (*request.Request, *CreateDBInstanceOutput)

	CreateDatabaseCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDatabaseCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDatabaseCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDatabase(*CreateDatabaseInput) (*CreateDatabaseOutput, error)
	CreateDatabaseWithContext(volcengine.Context, *CreateDatabaseInput, ...request.Option) (*CreateDatabaseOutput, error)
	CreateDatabaseRequest(*CreateDatabaseInput) (*request.Request, *CreateDatabaseOutput)

	DeleteAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAllowList(*DeleteAllowListInput) (*DeleteAllowListOutput, error)
	DeleteAllowListWithContext(volcengine.Context, *DeleteAllowListInput, ...request.Option) (*DeleteAllowListOutput, error)
	DeleteAllowListRequest(*DeleteAllowListInput) (*request.Request, *DeleteAllowListOutput)

	DeleteDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBAccount(*DeleteDBAccountInput) (*DeleteDBAccountOutput, error)
	DeleteDBAccountWithContext(volcengine.Context, *DeleteDBAccountInput, ...request.Option) (*DeleteDBAccountOutput, error)
	DeleteDBAccountRequest(*DeleteDBAccountInput) (*request.Request, *DeleteDBAccountOutput)

	DeleteDBEndpointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBEndpointCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBEndpointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBEndpoint(*DeleteDBEndpointInput) (*DeleteDBEndpointOutput, error)
	DeleteDBEndpointWithContext(volcengine.Context, *DeleteDBEndpointInput, ...request.Option) (*DeleteDBEndpointOutput, error)
	DeleteDBEndpointRequest(*DeleteDBEndpointInput) (*request.Request, *DeleteDBEndpointOutput)

	DeleteDBEndpointPublicAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBEndpointPublicAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBEndpointPublicAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBEndpointPublicAddress(*DeleteDBEndpointPublicAddressInput) (*DeleteDBEndpointPublicAddressOutput, error)
	DeleteDBEndpointPublicAddressWithContext(volcengine.Context, *DeleteDBEndpointPublicAddressInput, ...request.Option) (*DeleteDBEndpointPublicAddressOutput, error)
	DeleteDBEndpointPublicAddressRequest(*DeleteDBEndpointPublicAddressInput) (*request.Request, *DeleteDBEndpointPublicAddressOutput)

	DeleteDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBInstance(*DeleteDBInstanceInput) (*DeleteDBInstanceOutput, error)
	DeleteDBInstanceWithContext(volcengine.Context, *DeleteDBInstanceInput, ...request.Option) (*DeleteDBInstanceOutput, error)
	DeleteDBInstanceRequest(*DeleteDBInstanceInput) (*request.Request, *DeleteDBInstanceOutput)

	DeleteDatabaseCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDatabaseCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDatabaseCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDatabase(*DeleteDatabaseInput) (*DeleteDatabaseOutput, error)
	DeleteDatabaseWithContext(volcengine.Context, *DeleteDatabaseInput, ...request.Option) (*DeleteDatabaseOutput, error)
	DeleteDatabaseRequest(*DeleteDatabaseInput) (*request.Request, *DeleteDatabaseOutput)

	DescribeAllowListDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAllowListDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAllowListDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAllowListDetail(*DescribeAllowListDetailInput) (*DescribeAllowListDetailOutput, error)
	DescribeAllowListDetailWithContext(volcengine.Context, *DescribeAllowListDetailInput, ...request.Option) (*DescribeAllowListDetailOutput, error)
	DescribeAllowListDetailRequest(*DescribeAllowListDetailInput) (*request.Request, *DescribeAllowListDetailOutput)

	DescribeAllowListsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAllowListsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAllowListsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAllowLists(*DescribeAllowListsInput) (*DescribeAllowListsOutput, error)
	DescribeAllowListsWithContext(volcengine.Context, *DescribeAllowListsInput, ...request.Option) (*DescribeAllowListsOutput, error)
	DescribeAllowListsRequest(*DescribeAllowListsInput) (*request.Request, *DescribeAllowListsOutput)

	DescribeAvailabilityZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*DescribeAvailabilityZonesInput) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesWithContext(volcengine.Context, *DescribeAvailabilityZonesInput, ...request.Option) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesRequest(*DescribeAvailabilityZonesInput) (*request.Request, *DescribeAvailabilityZonesOutput)

	DescribeBackupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackups(*DescribeBackupsInput) (*DescribeBackupsOutput, error)
	DescribeBackupsWithContext(volcengine.Context, *DescribeBackupsInput, ...request.Option) (*DescribeBackupsOutput, error)
	DescribeBackupsRequest(*DescribeBackupsInput) (*request.Request, *DescribeBackupsOutput)

	DescribeDBAccountsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBAccountsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBAccountsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBAccounts(*DescribeDBAccountsInput) (*DescribeDBAccountsOutput, error)
	DescribeDBAccountsWithContext(volcengine.Context, *DescribeDBAccountsInput, ...request.Option) (*DescribeDBAccountsOutput, error)
	DescribeDBAccountsRequest(*DescribeDBAccountsInput) (*request.Request, *DescribeDBAccountsOutput)

	DescribeDBInstanceDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceDetail(*DescribeDBInstanceDetailInput) (*DescribeDBInstanceDetailOutput, error)
	DescribeDBInstanceDetailWithContext(volcengine.Context, *DescribeDBInstanceDetailInput, ...request.Option) (*DescribeDBInstanceDetailOutput, error)
	DescribeDBInstanceDetailRequest(*DescribeDBInstanceDetailInput) (*request.Request, *DescribeDBInstanceDetailOutput)

	DescribeDBInstanceParametersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceParametersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceParametersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceParameters(*DescribeDBInstanceParametersInput) (*DescribeDBInstanceParametersOutput, error)
	DescribeDBInstanceParametersWithContext(volcengine.Context, *DescribeDBInstanceParametersInput, ...request.Option) (*DescribeDBInstanceParametersOutput, error)
	DescribeDBInstanceParametersRequest(*DescribeDBInstanceParametersInput) (*request.Request, *DescribeDBInstanceParametersOutput)

	DescribeDBInstanceParametersLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceParametersLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceParametersLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceParametersLog(*DescribeDBInstanceParametersLogInput) (*DescribeDBInstanceParametersLogOutput, error)
	DescribeDBInstanceParametersLogWithContext(volcengine.Context, *DescribeDBInstanceParametersLogInput, ...request.Option) (*DescribeDBInstanceParametersLogOutput, error)
	DescribeDBInstanceParametersLogRequest(*DescribeDBInstanceParametersLogInput) (*request.Request, *DescribeDBInstanceParametersLogOutput)

	DescribeDBInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstances(*DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error)
	DescribeDBInstancesWithContext(volcengine.Context, *DescribeDBInstancesInput, ...request.Option) (*DescribeDBInstancesOutput, error)
	DescribeDBInstancesRequest(*DescribeDBInstancesInput) (*request.Request, *DescribeDBInstancesOutput)

	DescribeDatabasesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDatabasesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDatabasesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDatabases(*DescribeDatabasesInput) (*DescribeDatabasesOutput, error)
	DescribeDatabasesWithContext(volcengine.Context, *DescribeDatabasesInput, ...request.Option) (*DescribeDatabasesOutput, error)
	DescribeDatabasesRequest(*DescribeDatabasesInput) (*request.Request, *DescribeDatabasesOutput)

	DescribeRecoverableTimeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRecoverableTimeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRecoverableTimeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRecoverableTime(*DescribeRecoverableTimeInput) (*DescribeRecoverableTimeOutput, error)
	DescribeRecoverableTimeWithContext(volcengine.Context, *DescribeRecoverableTimeInput, ...request.Option) (*DescribeRecoverableTimeOutput, error)
	DescribeRecoverableTimeRequest(*DescribeRecoverableTimeInput) (*request.Request, *DescribeRecoverableTimeOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(volcengine.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DisassociateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateAllowList(*DisassociateAllowListInput) (*DisassociateAllowListOutput, error)
	DisassociateAllowListWithContext(volcengine.Context, *DisassociateAllowListInput, ...request.Option) (*DisassociateAllowListOutput, error)
	DisassociateAllowListRequest(*DisassociateAllowListInput) (*request.Request, *DisassociateAllowListOutput)

	DownloadBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DownloadBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DownloadBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DownloadBackup(*DownloadBackupInput) (*DownloadBackupOutput, error)
	DownloadBackupWithContext(volcengine.Context, *DownloadBackupInput, ...request.Option) (*DownloadBackupOutput, error)
	DownloadBackupRequest(*DownloadBackupInput) (*request.Request, *DownloadBackupOutput)

	GetBackupDownloadLinkCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetBackupDownloadLinkCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetBackupDownloadLinkCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetBackupDownloadLink(*GetBackupDownloadLinkInput) (*GetBackupDownloadLinkOutput, error)
	GetBackupDownloadLinkWithContext(volcengine.Context, *GetBackupDownloadLinkInput, ...request.Option) (*GetBackupDownloadLinkOutput, error)
	GetBackupDownloadLinkRequest(*GetBackupDownloadLinkInput) (*request.Request, *GetBackupDownloadLinkOutput)

	GrantDBAccountPrivilegeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GrantDBAccountPrivilegeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GrantDBAccountPrivilegeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GrantDBAccountPrivilege(*GrantDBAccountPrivilegeInput) (*GrantDBAccountPrivilegeOutput, error)
	GrantDBAccountPrivilegeWithContext(volcengine.Context, *GrantDBAccountPrivilegeInput, ...request.Option) (*GrantDBAccountPrivilegeOutput, error)
	GrantDBAccountPrivilegeRequest(*GrantDBAccountPrivilegeInput) (*request.Request, *GrantDBAccountPrivilegeOutput)

	GrantDatabasePrivilegeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GrantDatabasePrivilegeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GrantDatabasePrivilegeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GrantDatabasePrivilege(*GrantDatabasePrivilegeInput) (*GrantDatabasePrivilegeOutput, error)
	GrantDatabasePrivilegeWithContext(volcengine.Context, *GrantDatabasePrivilegeInput, ...request.Option) (*GrantDatabasePrivilegeOutput, error)
	GrantDatabasePrivilegeRequest(*GrantDatabasePrivilegeInput) (*request.Request, *GrantDatabasePrivilegeOutput)

	ModifyAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAllowList(*ModifyAllowListInput) (*ModifyAllowListOutput, error)
	ModifyAllowListWithContext(volcengine.Context, *ModifyAllowListInput, ...request.Option) (*ModifyAllowListOutput, error)
	ModifyAllowListRequest(*ModifyAllowListInput) (*request.Request, *ModifyAllowListOutput)

	ModifyDBEndpointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBEndpointCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBEndpointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBEndpoint(*ModifyDBEndpointInput) (*ModifyDBEndpointOutput, error)
	ModifyDBEndpointWithContext(volcengine.Context, *ModifyDBEndpointInput, ...request.Option) (*ModifyDBEndpointOutput, error)
	ModifyDBEndpointRequest(*ModifyDBEndpointInput) (*request.Request, *ModifyDBEndpointOutput)

	ModifyDBEndpointAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBEndpointAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBEndpointAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBEndpointAddress(*ModifyDBEndpointAddressInput) (*ModifyDBEndpointAddressOutput, error)
	ModifyDBEndpointAddressWithContext(volcengine.Context, *ModifyDBEndpointAddressInput, ...request.Option) (*ModifyDBEndpointAddressOutput, error)
	ModifyDBEndpointAddressRequest(*ModifyDBEndpointAddressInput) (*request.Request, *ModifyDBEndpointAddressOutput)

	ModifyDBEndpointDNSCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBEndpointDNSCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBEndpointDNSCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBEndpointDNS(*ModifyDBEndpointDNSInput) (*ModifyDBEndpointDNSOutput, error)
	ModifyDBEndpointDNSWithContext(volcengine.Context, *ModifyDBEndpointDNSInput, ...request.Option) (*ModifyDBEndpointDNSOutput, error)
	ModifyDBEndpointDNSRequest(*ModifyDBEndpointDNSInput) (*request.Request, *ModifyDBEndpointDNSOutput)

	ModifyDBInstanceParametersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceParametersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceParametersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceParameters(*ModifyDBInstanceParametersInput) (*ModifyDBInstanceParametersOutput, error)
	ModifyDBInstanceParametersWithContext(volcengine.Context, *ModifyDBInstanceParametersInput, ...request.Option) (*ModifyDBInstanceParametersOutput, error)
	ModifyDBInstanceParametersRequest(*ModifyDBInstanceParametersInput) (*request.Request, *ModifyDBInstanceParametersOutput)

	ModifyDBInstanceSpecCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceSpecCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceSpecCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceSpec(*ModifyDBInstanceSpecInput) (*ModifyDBInstanceSpecOutput, error)
	ModifyDBInstanceSpecWithContext(volcengine.Context, *ModifyDBInstanceSpecInput, ...request.Option) (*ModifyDBInstanceSpecOutput, error)
	ModifyDBInstanceSpecRequest(*ModifyDBInstanceSpecInput) (*request.Request, *ModifyDBInstanceSpecOutput)

	ResetDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ResetDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetDBAccount(*ResetDBAccountInput) (*ResetDBAccountOutput, error)
	ResetDBAccountWithContext(volcengine.Context, *ResetDBAccountInput, ...request.Option) (*ResetDBAccountOutput, error)
	ResetDBAccountRequest(*ResetDBAccountInput) (*request.Request, *ResetDBAccountOutput)

	RestartDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartDBInstance(*RestartDBInstanceInput) (*RestartDBInstanceOutput, error)
	RestartDBInstanceWithContext(volcengine.Context, *RestartDBInstanceInput, ...request.Option) (*RestartDBInstanceOutput, error)
	RestartDBInstanceRequest(*RestartDBInstanceInput) (*request.Request, *RestartDBInstanceOutput)

	RestoreToNewInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreToNewInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreToNewInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreToNewInstance(*RestoreToNewInstanceInput) (*RestoreToNewInstanceOutput, error)
	RestoreToNewInstanceWithContext(volcengine.Context, *RestoreToNewInstanceInput, ...request.Option) (*RestoreToNewInstanceOutput, error)
	RestoreToNewInstanceRequest(*RestoreToNewInstanceInput) (*request.Request, *RestoreToNewInstanceOutput)

	RevokeDBAccountPrivilegeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeDBAccountPrivilegeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeDBAccountPrivilegeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeDBAccountPrivilege(*RevokeDBAccountPrivilegeInput) (*RevokeDBAccountPrivilegeOutput, error)
	RevokeDBAccountPrivilegeWithContext(volcengine.Context, *RevokeDBAccountPrivilegeInput, ...request.Option) (*RevokeDBAccountPrivilegeOutput, error)
	RevokeDBAccountPrivilegeRequest(*RevokeDBAccountPrivilegeInput) (*request.Request, *RevokeDBAccountPrivilegeOutput)

	RevokeDatabasePrivilegeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeDatabasePrivilegeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeDatabasePrivilegeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeDatabasePrivilege(*RevokeDatabasePrivilegeInput) (*RevokeDatabasePrivilegeOutput, error)
	RevokeDatabasePrivilegeWithContext(volcengine.Context, *RevokeDatabasePrivilegeInput, ...request.Option) (*RevokeDatabasePrivilegeOutput, error)
	RevokeDatabasePrivilegeRequest(*RevokeDatabasePrivilegeInput) (*request.Request, *RevokeDatabasePrivilegeOutput)

	UpgradeAllowListVersionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpgradeAllowListVersionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpgradeAllowListVersionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpgradeAllowListVersion(*UpgradeAllowListVersionInput) (*UpgradeAllowListVersionOutput, error)
	UpgradeAllowListVersionWithContext(volcengine.Context, *UpgradeAllowListVersionInput, ...request.Option) (*UpgradeAllowListVersionOutput, error)
	UpgradeAllowListVersionRequest(*UpgradeAllowListVersionInput) (*request.Request, *UpgradeAllowListVersionOutput)
}

var _ RDSMYSQLV2API = (*RDSMYSQLV2)(nil)
