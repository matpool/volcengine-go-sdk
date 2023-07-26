// Example Code generated by Beijing Volcanoengine Technology.
package cenexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/cen"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeCenAttachedInstances() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	describeCenAttachedInstancesInput := &cen.DescribeCenAttachedInstancesInput{
		CenId:      volcengine.String("cen-7qthudw0ll6jmc****"),
		PageNumber: volcengine.Int64(1),
		PageSize:   volcengine.Int64(20),
	}

	resp, err := svc.DescribeCenAttachedInstances(describeCenAttachedInstancesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
