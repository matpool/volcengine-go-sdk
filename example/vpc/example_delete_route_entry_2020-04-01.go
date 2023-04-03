package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteRouteEntry() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	deleteRouteEntryInput := &vpc.DeleteRouteEntryInput{
		RouteEntryId: volcengine.String("rte-3tj9gw2pwq3vj0x****"),
	}

	resp, err := svc.DeleteRouteEntry(deleteRouteEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
