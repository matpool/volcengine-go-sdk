package ecsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeInstanceVncUrl() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	describeInstanceVncUrlInput := &ecs.DescribeInstanceVncUrlInput{
		InstanceId: volcengine.String("i-l8u10sauiu9qj0h*****"),
	}

	resp, err := svc.DescribeInstanceVncUrl(describeInstanceVncUrlInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
