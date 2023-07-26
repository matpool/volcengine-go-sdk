// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/autoscaling"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DisableScalingPolicy() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	disableScalingPolicyInput := &autoscaling.DisableScalingPolicyInput{
		ScalingPolicyId: volcengine.String("sp-ybmvamf8uql8j1fl****"),
	}

	resp, err := svc.DisableScalingPolicy(disableScalingPolicyInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
