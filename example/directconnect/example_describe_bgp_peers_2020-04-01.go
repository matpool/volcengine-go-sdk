// Example Code generated by Beijing Volcanoengine Technology.
package directconnectexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/directconnect"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeBgpPeers() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	describeBgpPeersInput := &directconnect.DescribeBgpPeersInput{
		BgpPeerIds: volcengine.StringSlice([]string{"bgp-2752hz4teko3k7f4c****"}),
		PageNumber: volcengine.Int64(1),
		PageSize:   volcengine.Int64(20),
	}

	resp, err := svc.DescribeBgpPeers(describeBgpPeersInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
