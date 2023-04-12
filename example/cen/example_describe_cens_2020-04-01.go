package cenexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeCens() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	describeCensInput := &cen.DescribeCensInput{
		PageNumber: volcengine.Int64(1),
		PageSize:   volcengine.Int64(20),
	}

	resp, err := svc.DescribeCens(describeCensInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
