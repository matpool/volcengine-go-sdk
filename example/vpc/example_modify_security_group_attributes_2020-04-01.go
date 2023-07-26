// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifySecurityGroupAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifySecurityGroupAttributesInput := &vpc.ModifySecurityGroupAttributesInput{
		SecurityGroupId:   volcengine.String("sg-bp67acfmxazb4p****"),
		SecurityGroupName: volcengine.String("sg-1"),
	}

	resp, err := svc.ModifySecurityGroupAttributes(modifySecurityGroupAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
