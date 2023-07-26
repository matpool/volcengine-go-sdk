// Example Code generated by Beijing Volcanoengine Technology.
package cenexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/cen"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func DescribeCenRouteEntries() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	describeCenRouteEntriesInput := &cen.DescribeCenRouteEntriesInput{
		PageNumber: volcengine.Int64(1),
		PageSize:   volcengine.Int64(20),
	}

	resp, err := svc.DescribeCenRouteEntries(describeCenRouteEntriesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
