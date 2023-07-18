// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/autoscaling"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeScalingConfigurations() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	describeScalingConfigurationsInput := &autoscaling.DescribeScalingConfigurationsInput{
		ScalingGroupId: volcengine.String("scg-ybmssdnnhn5pkgyd****"),
	}

	resp, err := svc.DescribeScalingConfigurations(describeScalingConfigurationsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
