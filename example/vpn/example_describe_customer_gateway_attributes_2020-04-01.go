package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeCustomerGatewayAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeCustomerGatewayAttributesInput := &vpn.DescribeCustomerGatewayAttributesInput{
		CustomerGatewayId: volcengine.String("cgw-7qthudw0ll6jmc****"),
	}

	resp, err := svc.DescribeCustomerGatewayAttributes(describeCustomerGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
