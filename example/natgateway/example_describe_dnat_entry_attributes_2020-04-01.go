// Example Code generated by Beijing Volcanoengine Technology.
package natgatewayexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/natgateway"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeDnatEntryAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	describeDnatEntryAttributesInput := &natgateway.DescribeDnatEntryAttributesInput{
		DnatEntryId: volcengine.String("dnat-342abc3bc3****"),
	}

	resp, err := svc.DescribeDnatEntryAttributes(describeDnatEntryAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
