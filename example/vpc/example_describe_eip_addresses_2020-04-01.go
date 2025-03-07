// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeEipAddresses() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	describeEipAddressesInput := &vpc.DescribeEipAddressesInput{
		BillingType: volcengine.Int64(2),
		PageNumber:  volcengine.Int64(1),
		PageSize:    volcengine.Int64(20),
	}

	resp, err := svc.DescribeEipAddresses(describeEipAddressesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
