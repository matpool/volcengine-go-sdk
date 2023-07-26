package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeVpnConnections() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeVpnConnectionsInput := &vpn.DescribeVpnConnectionsInput{
		PageNumber:       volcengine.Int64(1),
		PageSize:         volcengine.Int64(20),
		VpnConnectionIds: volcengine.StringSlice([]string{"vgc-2bzvqi8kerd342dx0eg2f****"}),
	}

	resp, err := svc.DescribeVpnConnections(describeVpnConnectionsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
