// Example Code generated by Beijing Volcanoengine Technology.
package ecsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/ecs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeImageSharePermission() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	describeImageSharePermissionInput := &ecs.DescribeImageSharePermissionInput{
		ImageId: volcengine.String("image-ebgy1og99tfct0gw****"),
	}

	resp, err := svc.DescribeImageSharePermission(describeImageSharePermissionInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
