// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/autoscaling"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyLifecycleHook() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	modifyLifecycleHookInput := &autoscaling.ModifyLifecycleHookInput{
		LifecycleHookId:      volcengine.String("sgh-ybrzhc5ht08hc******"),
		LifecycleHookPolicy:  volcengine.String("CONTINUE"),
		LifecycleHookTimeout: volcengine.Int32(30),
		LifecycleHookType:    volcengine.String("SCALE_IN"),
	}

	resp, err := svc.ModifyLifecycleHook(modifyLifecycleHookInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
