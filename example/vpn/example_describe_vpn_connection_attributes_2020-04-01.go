package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeVpnConnectionAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeVpnConnectionAttributesInput := &vpn.DescribeVpnConnectionAttributesInput{
		VpnConnectionId: volcengine.String("vgc-2bzvqi8kerd342dx0eg2f****"),
	}

	resp, err := svc.DescribeVpnConnectionAttributes(describeVpnConnectionAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
