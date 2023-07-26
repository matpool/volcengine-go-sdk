// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func SetLoadBalancerRenewal() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	setLoadBalancerRenewalInput := &clb.SetLoadBalancerRenewalInput{
		LoadBalancerId: volcengine.String("clb-bp1b6c719dfa08ex****"),
		RenewType:      volcengine.Int64(2),
	}

	resp, err := svc.SetLoadBalancerRenewal(setLoadBalancerRenewalInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
