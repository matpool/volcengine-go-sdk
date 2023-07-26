package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func CreateVpnGatewayRoute() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	createVpnGatewayRouteInput := &vpn.CreateVpnGatewayRouteInput{
		DestinationCidrBlock: volcengine.String("172.XX.XX.0/24"),
		NextHopId:            volcengine.String("vgc-7qthudw0ll6jmc****"),
		VpnGatewayId:         volcengine.String("vgw-12bfa2du7fojk17q7y1rk****"),
	}

	resp, err := svc.CreateVpnGatewayRoute(createVpnGatewayRouteInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
