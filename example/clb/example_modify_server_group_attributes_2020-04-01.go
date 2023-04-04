package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyServerGroupAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	reqServers0 := &clb.ServerForModifyServerGroupAttributesInput{
		Port:     volcengine.Int64(88),
		ServerId: volcengine.String("rs-mjc9b2p0v6rk5smt1b27****"),
		Weight:   volcengine.Int64(100),
	}
	modifyServerGroupAttributesInput := &clb.ModifyServerGroupAttributesInput{
		ServerGroupId:   volcengine.String("rsp-bp1o94dp5i6ea****"),
		ServerGroupName: volcengine.String("myservergroup"),
		Servers:         []*clb.ServerForModifyServerGroupAttributesInput{reqServers0},
	}

	resp, err := svc.ModifyServerGroupAttributes(modifyServerGroupAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
