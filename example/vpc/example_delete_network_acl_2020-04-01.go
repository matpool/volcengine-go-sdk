// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DeleteNetworkAcl() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	deleteNetworkAclInput := &vpc.DeleteNetworkAclInput{
		NetworkAclId: volcengine.String("acl-bp1fg655nh68xyz9****"),
	}

	resp, err := svc.DeleteNetworkAcl(deleteNetworkAclInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
