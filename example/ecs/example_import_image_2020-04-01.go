// Example Code generated by Beijing Volcanoengine Technology.
package ecsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/ecs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ImportImage() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	importImageInput := &ecs.ImportImageInput{
		Architecture:    volcengine.String("amd64"),
		ImageName:       volcengine.String("image-1"),
		OsType:          volcengine.String("Linux"),
		Platform:        volcengine.String("CentOS"),
		PlatformVersion: volcengine.String("7.6"),
		Url:             volcengine.String("xxx"),
	}

	resp, err := svc.ImportImage(importImageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
