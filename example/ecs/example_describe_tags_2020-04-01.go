package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeTags() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	reqTagFilters0 := &ecs.TagFilterForDescribeTagsInput{
		Key:    volcengine.String("k1"),
		Values: volcengine.StringSlice([]string{"v1", "v2"}),
	}
	describeTagsInput := &ecs.DescribeTagsInput{
		ResourceIds:  volcengine.StringSlice([]string{"i-l8u0p77yseabkpak****", "i-l8u0p7xyseabkbak****"}),
		ResourceType: volcengine.String("instance"),
		TagFilters:   []*ecs.TagFilterForDescribeTagsInput{reqTagFilters0},
	}

	resp, err := svc.DescribeTags(describeTagsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
