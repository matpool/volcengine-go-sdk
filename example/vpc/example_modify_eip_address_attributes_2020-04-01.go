// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyEipAddressAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifyEipAddressAttributesInput := &vpc.ModifyEipAddressAttributesInput{
		AllocationId: volcengine.String("eip-2zb7uj8xscd****"),
		Bandwidth:    volcengine.Int64(10),
	}

	resp, err := svc.ModifyEipAddressAttributes(modifyEipAddressAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
