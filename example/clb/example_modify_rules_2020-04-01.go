// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/clb"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func ModifyRules() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	reqRules := &clb.RuleForModifyRulesInput{
		Description:   volcengine.String("test"),
		RuleId:        volcengine.String("rule-2fegss1cplxxc5oxruvvq****"),
		ServerGroupId: volcengine.String("rsp-bp1o94dp5i6ea****"),
	}
	modifyRulesInput := &clb.ModifyRulesInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
		Rules:      []*clb.RuleForModifyRulesInput{reqRules},
	}

	resp, err := svc.ModifyRules(modifyRulesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
