// Example Code generated by Beijing Volcanoengine Technology.
package vpcexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpc"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifySecurityGroupRuleDescriptionsIngress() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifySecurityGroupRuleDescriptionsIngressInput := &vpc.ModifySecurityGroupRuleDescriptionsIngressInput{
		CidrIp:          volcengine.String("10.XX.XX.0/8"),
		Description:     volcengine.String("S_G_R_I_Description"),
		Policy:          volcengine.String("accept"),
		PortEnd:         volcengine.Int64(22),
		PortStart:       volcengine.Int64(22),
		Priority:        volcengine.Int64(1),
		Protocol:        volcengine.String("tcp"),
		SecurityGroupId: volcengine.String("sg-bp67acfmxazb4p****"),
	}

	resp, err := svc.ModifySecurityGroupRuleDescriptionsIngress(modifySecurityGroupRuleDescriptionsIngressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
