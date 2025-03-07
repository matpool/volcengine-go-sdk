// Example Code generated by Beijing Volcanoengine Technology.
package directconnectexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/directconnect"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DeleteDirectConnectGatewayRoute() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	deleteDirectConnectGatewayRouteInput := &directconnect.DeleteDirectConnectGatewayRouteInput{
		DirectConnectGatewayRouteId: volcengine.String("dcr-7qthudw0ll6jmc****"),
	}

	resp, err := svc.DeleteDirectConnectGatewayRoute(deleteDirectConnectGatewayRouteInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
