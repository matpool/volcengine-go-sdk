// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/autoscaling"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DeleteLifecycleHook() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	deleteLifecycleHookInput := &autoscaling.DeleteLifecycleHookInput{
		LifecycleHookId: volcengine.String("sgh-ybrzhc5ht08hccn*****"),
	}

	resp, err := svc.DeleteLifecycleHook(deleteLifecycleHookInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
