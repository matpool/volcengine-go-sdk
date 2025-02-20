// Example Code generated by Beijing Volcanoengine Technology.
package cenexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/cen"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyCenAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	modifyCenAttributesInput := &cen.ModifyCenAttributesInput{
		CenId:       volcengine.String("cen-7qthudw0ll6jmc****"),
		CenName:     volcengine.String("cen-test"),
		Description: volcengine.String("test"),
	}

	resp, err := svc.ModifyCenAttributes(modifyCenAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
