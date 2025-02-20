// Example Code generated by Beijing Volcanoengine Technology.
package ecsexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/ecs"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func AssociateInstancesIamRole() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	associateInstancesIamRoleInput := &ecs.AssociateInstancesIamRoleInput{
		IamRoleName: volcengine.String("EcsTestRole"),
		InstanceIds: volcengine.StringSlice([]string{"i-3tiefmkskq3vj0******"}),
	}

	resp, err := svc.AssociateInstancesIamRole(associateInstancesIamRoleInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
