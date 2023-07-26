package vpnexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vpn"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyCustomerGatewayAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	modifyCustomerGatewayAttributesInput := &vpn.ModifyCustomerGatewayAttributesInput{
		CustomerGatewayId:   volcengine.String("cgw-2d670j2o9lc0058ozfddg****"),
		CustomerGatewayName: volcengine.String("test"),
		Description:         volcengine.String("test"),
	}

	resp, err := svc.ModifyCustomerGatewayAttributes(modifyCustomerGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
