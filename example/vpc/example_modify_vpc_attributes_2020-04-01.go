// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyVpcAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifyVpcAttributesInput := &vpc.ModifyVpcAttributesInput{
		VpcId:   volcengine.String("vpc-bp15zct37pq72zv****"),
		VpcName: volcengine.String("vpc-1"),
	}

	resp, err := svc.ModifyVpcAttributes(modifyVpcAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
