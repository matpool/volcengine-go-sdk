// Example Code generated by Beijing Volcanoengine Technology.
package cenexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/cen"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeCenBandwidthPackageAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	describeCenBandwidthPackageAttributesInput := &cen.DescribeCenBandwidthPackageAttributesInput{
		CenBandwidthPackageId: volcengine.String("cbp-bp1o94dp5****"),
	}

	resp, err := svc.DescribeCenBandwidthPackageAttributes(describeCenBandwidthPackageAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
