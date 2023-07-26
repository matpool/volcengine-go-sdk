// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeLoadBalancers() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeLoadBalancersInput := &clb.DescribeLoadBalancersInput{
		PageNumber: volcengine.Int64(1),
		PageSize:   volcengine.Int64(10),
		VpcId:      volcengine.String("vpc-13fd2oy7dsiyo3n6nu4ye****"),
	}

	resp, err := svc.DescribeLoadBalancers(describeLoadBalancersInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
