// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeListenerHealth() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeListenerHealthInput := &clb.DescribeListenerHealthInput{
		ListenerId:    volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
		OnlyUnHealthy: volcengine.String("true"),
		PageNumber:    volcengine.Int64(1),
		PageSize:      volcengine.Int64(20),
	}

	resp, err := svc.DescribeListenerHealth(describeListenerHealthInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
