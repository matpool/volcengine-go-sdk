// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func AllocateEipAddress() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	allocateEipAddressInput := &vpc.AllocateEipAddressInput{
		Bandwidth:   volcengine.Int64(10),
		BillingType: volcengine.Int64(2),
		ISP:         volcengine.String("BGP"),
		Name:        volcengine.String("eip-1"),
	}

	resp, err := svc.AllocateEipAddress(allocateEipAddressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
