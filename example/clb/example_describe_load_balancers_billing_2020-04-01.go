// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeLoadBalancersBilling() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeLoadBalancersBillingInput := &clb.DescribeLoadBalancersBillingInput{
		LoadBalancerIds: volcengine.StringSlice([]string{"clb-bp1b6c719dfa08ex****"}),
	}

	resp, err := svc.DescribeLoadBalancersBilling(describeLoadBalancersBillingInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
