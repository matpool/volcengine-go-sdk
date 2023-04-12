package ecsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifySubscriptionEventTypes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	modifySubscriptionEventTypesInput := &ecs.ModifySubscriptionEventTypesInput{
		EventTypes:     volcengine.StringSlice([]string{"SystemFailure.Stop:Succeeded", "SystemFailure.Stop:Succeeded"}),
		SubscriptionId: volcengine.String("s-6js1al1y9665lp******"),
	}

	resp, err := svc.ModifySubscriptionEventTypes(modifySubscriptionEventTypesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
