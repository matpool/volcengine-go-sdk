// Example Code generated by Beijing Volcanoengine Technology.
package cenexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/cen"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func AttachInstanceToCen() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	attachInstanceToCenInput := &cen.AttachInstanceToCenInput{
		CenId:            volcengine.String("cen-7qthudw0ll6jmc****"),
		InstanceId:       volcengine.String("vpc-2fexiqjlgjif45oxruvso****"),
		InstanceRegionId: volcengine.String("cn-beijing"),
		InstanceType:     volcengine.String("vpc"),
	}

	resp, err := svc.AttachInstanceToCen(attachInstanceToCenInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
