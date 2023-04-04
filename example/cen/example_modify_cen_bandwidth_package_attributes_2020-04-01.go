package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyCenBandwidthPackageAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	modifyCenBandwidthPackageAttributesInput := &cen.ModifyCenBandwidthPackageAttributesInput{
		CenBandwidthPackageId: volcengine.String("cbp-4c2zaavbvh5fx****"),
	}

	resp, err := svc.ModifyCenBandwidthPackageAttributes(modifyCenBandwidthPackageAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
