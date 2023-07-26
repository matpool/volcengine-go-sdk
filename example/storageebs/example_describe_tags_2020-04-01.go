// Example Code generated by Beijing Volcanoengine Technology.
package storageebsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/storageebs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeTags() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := storageebs.New(sess)
	describeTagsInput := &storageebs.DescribeTagsInput{
		ResourceIds:  volcengine.StringSlice([]string{"vol-76pwr6rilv8lzwxv*****", "vol-76pwr6rilv8lzwxv****"}),
		ResourceType: volcengine.String("volume"),
	}

	resp, err := svc.DescribeTags(describeTagsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
