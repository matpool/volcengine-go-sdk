// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeListenerAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeListenerAttributesInput := &clb.DescribeListenerAttributesInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
	}

	resp, err := svc.DescribeListenerAttributes(describeListenerAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
