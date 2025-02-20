// Example Code generated by Beijing Volcanoengine Technology.
package storageebsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/storageebs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeVolumes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := storageebs.New(sess)
	describeVolumesInput := &storageebs.DescribeVolumesInput{
		VolumeIds: volcengine.StringSlice([]string{"vol-3tirbh6lka3vj0x1****", "vol-3tj20xxpd63vj0wy****"}),
	}

	resp, err := svc.DescribeVolumes(describeVolumesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
