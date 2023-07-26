package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeVpnGatewayAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeVpnGatewayAttributesInput := &vpn.DescribeVpnGatewayAttributesInput{
		VpnGatewayId: volcengine.String("vgw-12bfa2du7fojk17q7y1rk****"),
	}

	resp, err := svc.DescribeVpnGatewayAttributes(describeVpnGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
