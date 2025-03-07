// Example Code generated by Beijing Volcanoengine Technology.
package directconnectexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/directconnect"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyDirectConnectGatewayAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	modifyDirectConnectGatewayAttributesInput := &directconnect.ModifyDirectConnectGatewayAttributesInput{
		Description:              volcengine.String("test"),
		DirectConnectGatewayId:   volcengine.String("dcg-7qthudw0ll6jmc****"),
		DirectConnectGatewayName: volcengine.String("test"),
	}

	resp, err := svc.ModifyDirectConnectGatewayAttributes(modifyDirectConnectGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
