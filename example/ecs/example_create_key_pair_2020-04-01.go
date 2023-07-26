// Example Code generated by Beijing Volcanoengine Technology.
package ecsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/ecs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func CreateKeyPair() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	createKeyPairInput := &ecs.CreateKeyPairInput{
		KeyPairName: volcengine.String("ssh_key_pair"),
	}

	resp, err := svc.CreateKeyPair(createKeyPairInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
