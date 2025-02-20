// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyServerGroupAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	reqServers := &clb.ServerForModifyServerGroupAttributesInput{
		Port:     volcengine.Int64(88),
		ServerId: volcengine.String("rs-mjc9b2p0v6rk5smt1b27****"),
		Weight:   volcengine.Int64(100),
	}
	modifyServerGroupAttributesInput := &clb.ModifyServerGroupAttributesInput{
		ServerGroupId:   volcengine.String("rsp-bp1o94dp5i6ea****"),
		ServerGroupName: volcengine.String("myservergroup"),
		Servers:         []*clb.ServerForModifyServerGroupAttributesInput{reqServers},
	}

	resp, err := svc.ModifyServerGroupAttributes(modifyServerGroupAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
