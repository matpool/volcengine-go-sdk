package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyEipAddressAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifyEipAddressAttributesInput := &vpc.ModifyEipAddressAttributesInput{
		AllocationId: volcengine.String("eip-2zb7uj8xscd****"),
		Bandwidth:    volcengine.Int64(10),
	}

	resp, err := svc.ModifyEipAddressAttributes(modifyEipAddressAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
