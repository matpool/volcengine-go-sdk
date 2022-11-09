// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package natgatewayiface provides an interface to enable mocking the NATGATEWAY service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package natgateway

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// NATGATEWAYAPI provides an interface to enable mocking the
// natgateway.NATGATEWAY service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // NATGATEWAY.
//    func myFunc(svc NATGATEWAYAPI) bool {
//        // Make svc.CreateDnatEntry request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := natgateway.New(sess)
//
//        myFunc(svc)
//    }
//
type NATGATEWAYAPI interface {
	CreateDnatEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDnatEntryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDnatEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDnatEntry(*CreateDnatEntryInput) (*CreateDnatEntryOutput, error)
	CreateDnatEntryWithContext(volcengine.Context, *CreateDnatEntryInput, ...request.Option) (*CreateDnatEntryOutput, error)
	CreateDnatEntryRequest(*CreateDnatEntryInput) (*request.Request, *CreateDnatEntryOutput)

	CreateNatGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNatGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNatGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNatGateway(*CreateNatGatewayInput) (*CreateNatGatewayOutput, error)
	CreateNatGatewayWithContext(volcengine.Context, *CreateNatGatewayInput, ...request.Option) (*CreateNatGatewayOutput, error)
	CreateNatGatewayRequest(*CreateNatGatewayInput) (*request.Request, *CreateNatGatewayOutput)

	CreateSnatEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSnatEntryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSnatEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSnatEntry(*CreateSnatEntryInput) (*CreateSnatEntryOutput, error)
	CreateSnatEntryWithContext(volcengine.Context, *CreateSnatEntryInput, ...request.Option) (*CreateSnatEntryOutput, error)
	CreateSnatEntryRequest(*CreateSnatEntryInput) (*request.Request, *CreateSnatEntryOutput)

	DeleteDnatEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDnatEntryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDnatEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDnatEntry(*DeleteDnatEntryInput) (*DeleteDnatEntryOutput, error)
	DeleteDnatEntryWithContext(volcengine.Context, *DeleteDnatEntryInput, ...request.Option) (*DeleteDnatEntryOutput, error)
	DeleteDnatEntryRequest(*DeleteDnatEntryInput) (*request.Request, *DeleteDnatEntryOutput)

	DeleteNatGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNatGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNatGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNatGateway(*DeleteNatGatewayInput) (*DeleteNatGatewayOutput, error)
	DeleteNatGatewayWithContext(volcengine.Context, *DeleteNatGatewayInput, ...request.Option) (*DeleteNatGatewayOutput, error)
	DeleteNatGatewayRequest(*DeleteNatGatewayInput) (*request.Request, *DeleteNatGatewayOutput)

	DeleteSnatEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSnatEntryCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSnatEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSnatEntry(*DeleteSnatEntryInput) (*DeleteSnatEntryOutput, error)
	DeleteSnatEntryWithContext(volcengine.Context, *DeleteSnatEntryInput, ...request.Option) (*DeleteSnatEntryOutput, error)
	DeleteSnatEntryRequest(*DeleteSnatEntryInput) (*request.Request, *DeleteSnatEntryOutput)

	DescribeDnatEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDnatEntriesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDnatEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDnatEntries(*DescribeDnatEntriesInput) (*DescribeDnatEntriesOutput, error)
	DescribeDnatEntriesWithContext(volcengine.Context, *DescribeDnatEntriesInput, ...request.Option) (*DescribeDnatEntriesOutput, error)
	DescribeDnatEntriesRequest(*DescribeDnatEntriesInput) (*request.Request, *DescribeDnatEntriesOutput)

	DescribeDnatEntryAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDnatEntryAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDnatEntryAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDnatEntryAttributes(*DescribeDnatEntryAttributesInput) (*DescribeDnatEntryAttributesOutput, error)
	DescribeDnatEntryAttributesWithContext(volcengine.Context, *DescribeDnatEntryAttributesInput, ...request.Option) (*DescribeDnatEntryAttributesOutput, error)
	DescribeDnatEntryAttributesRequest(*DescribeDnatEntryAttributesInput) (*request.Request, *DescribeDnatEntryAttributesOutput)

	DescribeNatGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNatGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNatGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNatGatewayAttributes(*DescribeNatGatewayAttributesInput) (*DescribeNatGatewayAttributesOutput, error)
	DescribeNatGatewayAttributesWithContext(volcengine.Context, *DescribeNatGatewayAttributesInput, ...request.Option) (*DescribeNatGatewayAttributesOutput, error)
	DescribeNatGatewayAttributesRequest(*DescribeNatGatewayAttributesInput) (*request.Request, *DescribeNatGatewayAttributesOutput)

	DescribeNatGatewaysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNatGatewaysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNatGatewaysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNatGateways(*DescribeNatGatewaysInput) (*DescribeNatGatewaysOutput, error)
	DescribeNatGatewaysWithContext(volcengine.Context, *DescribeNatGatewaysInput, ...request.Option) (*DescribeNatGatewaysOutput, error)
	DescribeNatGatewaysRequest(*DescribeNatGatewaysInput) (*request.Request, *DescribeNatGatewaysOutput)

	DescribeSnatEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSnatEntriesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSnatEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSnatEntries(*DescribeSnatEntriesInput) (*DescribeSnatEntriesOutput, error)
	DescribeSnatEntriesWithContext(volcengine.Context, *DescribeSnatEntriesInput, ...request.Option) (*DescribeSnatEntriesOutput, error)
	DescribeSnatEntriesRequest(*DescribeSnatEntriesInput) (*request.Request, *DescribeSnatEntriesOutput)

	DescribeSnatEntryAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSnatEntryAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSnatEntryAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSnatEntryAttributes(*DescribeSnatEntryAttributesInput) (*DescribeSnatEntryAttributesOutput, error)
	DescribeSnatEntryAttributesWithContext(volcengine.Context, *DescribeSnatEntryAttributesInput, ...request.Option) (*DescribeSnatEntryAttributesOutput, error)
	DescribeSnatEntryAttributesRequest(*DescribeSnatEntryAttributesInput) (*request.Request, *DescribeSnatEntryAttributesOutput)

	ListNatGatewayAvailableZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListNatGatewayAvailableZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListNatGatewayAvailableZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListNatGatewayAvailableZones(*ListNatGatewayAvailableZonesInput) (*ListNatGatewayAvailableZonesOutput, error)
	ListNatGatewayAvailableZonesWithContext(volcengine.Context, *ListNatGatewayAvailableZonesInput, ...request.Option) (*ListNatGatewayAvailableZonesOutput, error)
	ListNatGatewayAvailableZonesRequest(*ListNatGatewayAvailableZonesInput) (*request.Request, *ListNatGatewayAvailableZonesOutput)

	ModifyDnatEntryAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDnatEntryAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDnatEntryAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDnatEntryAttributes(*ModifyDnatEntryAttributesInput) (*ModifyDnatEntryAttributesOutput, error)
	ModifyDnatEntryAttributesWithContext(volcengine.Context, *ModifyDnatEntryAttributesInput, ...request.Option) (*ModifyDnatEntryAttributesOutput, error)
	ModifyDnatEntryAttributesRequest(*ModifyDnatEntryAttributesInput) (*request.Request, *ModifyDnatEntryAttributesOutput)

	ModifyNatGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNatGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNatGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNatGatewayAttributes(*ModifyNatGatewayAttributesInput) (*ModifyNatGatewayAttributesOutput, error)
	ModifyNatGatewayAttributesWithContext(volcengine.Context, *ModifyNatGatewayAttributesInput, ...request.Option) (*ModifyNatGatewayAttributesOutput, error)
	ModifyNatGatewayAttributesRequest(*ModifyNatGatewayAttributesInput) (*request.Request, *ModifyNatGatewayAttributesOutput)

	ModifySnatEntryAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySnatEntryAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySnatEntryAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySnatEntryAttributes(*ModifySnatEntryAttributesInput) (*ModifySnatEntryAttributesOutput, error)
	ModifySnatEntryAttributesWithContext(volcengine.Context, *ModifySnatEntryAttributesInput, ...request.Option) (*ModifySnatEntryAttributesOutput, error)
	ModifySnatEntryAttributesRequest(*ModifySnatEntryAttributesInput) (*request.Request, *ModifySnatEntryAttributesOutput)
}

var _ NATGATEWAYAPI = (*NATGATEWAY)(nil)
