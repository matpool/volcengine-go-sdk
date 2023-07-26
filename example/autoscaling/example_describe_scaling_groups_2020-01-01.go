// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/autoscaling"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeScalingGroups() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	describeScalingGroupsInput := &autoscaling.DescribeScalingGroupsInput{
		ScalingGroupIds: volcengine.StringSlice([]string{"scg-ybmssdnnhn5pkgyd****"}),
	}

	resp, err := svc.DescribeScalingGroups(describeScalingGroupsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
