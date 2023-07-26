// Example Code generated by Beijing Volcanoengine Technology.
package ecsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/ecs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func CreateSubscription() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	createSubscriptionInput := &ecs.CreateSubscriptionInput{
		EventTypes: volcengine.StringSlice([]string{"SystemFailure.Stop:Succeeded", "SystemFailure.Stop:Succeeded"}),
		Type:       volcengine.String("Notification"),
	}

	resp, err := svc.CreateSubscription(createSubscriptionInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
