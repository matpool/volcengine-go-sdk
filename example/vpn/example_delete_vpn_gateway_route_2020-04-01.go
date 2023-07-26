package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DeleteVpnGatewayRoute() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	deleteVpnGatewayRouteInput := &vpn.DeleteVpnGatewayRouteInput{
		VpnGatewayRouteId: volcengine.String("vgr-7qthudw0ll6jmc****"),
	}

	resp, err := svc.DeleteVpnGatewayRoute(deleteVpnGatewayRouteInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
