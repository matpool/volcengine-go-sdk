// Example Code generated by Beijing Volcanoengine Technology.
package vkeemample

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/service/vke"
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/credentials"
	"github.com/matpool/volcengine-go-sdk/volcengine/session"
)

func CreateNodePool() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vke.New(sess)
	reqAutoScaling := &vke.AutoScalingForCreateNodePoolInput{
		DesiredReplicas: volcengine.Int32(0),
		Enabled:         volcengine.Bool(true),
		MaxReplicas:     volcengine.Int32(10),
		MinReplicas:     volcengine.Int32(0),
		Priority:        volcengine.Int32(10),
		SubnetPolicy:    volcengine.String("ZoneBalance"),
	}
	reqLabels0 := &vke.LabelForCreateNodePoolInput{
		Key:   volcengine.String("label-key"),
		Value: volcengine.String("label-value"),
	}
	reqTaints0 := &vke.TaintForCreateNodePoolInput{
		Effect: volcengine.String("NoSchedule"),
		Key:    volcengine.String("taint-key"),
		Value:  volcengine.String("taint-value"),
	}
	reqKubernetesConfig := &vke.KubernetesConfigForCreateNodePoolInput{
		Cordon: volcengine.Bool(false),
		Labels: []*vke.LabelForCreateNodePoolInput{reqLabels0},
		Taints: []*vke.TaintForCreateNodePoolInput{reqTaints0},
	}
	reqDataVolumes0 := &vke.DataVolumeForCreateNodePoolInput{
		Size: volcengine.Int32(20),
		Type: volcengine.String("ESSD_PL0"),
	}
	reqLogin := &vke.LoginForCreateNodePoolInput{
		Password: volcengine.String("UHdkMTIz***"),
	}
	reqSecurity := &vke.SecurityForCreateNodePoolInput{
		Login:              reqLogin,
		SecurityGroupIds:   volcengine.StringSlice([]string{"sg-2byy13cnsczy****"}),
		SecurityStrategies: volcengine.StringSlice([]string{"Hids"}),
	}
	reqSystemVolume := &vke.SystemVolumeForCreateNodePoolInput{
		Size: volcengine.Int32(40),
		Type: volcengine.String("ESSD_PL0"),
	}
	reqTags0 := &vke.TagForCreateNodePoolInput{
		Key:   volcengine.String("key"),
		Value: volcengine.String("value"),
	}
	reqNodeConfig := &vke.NodeConfigForCreateNodePoolInput{
		AdditionalContainerStorageEnabled: volcengine.Bool(true),
		DataVolumes:                       []*vke.DataVolumeForCreateNodePoolInput{reqDataVolumes0},
		InitializeScript:                  volcengine.String("ZWNobyAidG******"),
		InstanceTypeIds:                   volcengine.StringSlice([]string{"ecs.g1.xlarge"}),
		NamePrefix:                        volcengine.String("name-prefix"),
		Security:                          reqSecurity,
		SubnetIds:                         volcengine.StringSlice([]string{"subnet-3rf6vwbgkg****"}),
		SystemVolume:                      reqSystemVolume,
		Tags:                              []*vke.TagForCreateNodePoolInput{reqTags0},
	}
	reqTags1 := &vke.TagForCreateNodePoolInput{
		Key:   volcengine.String("key"),
		Value: volcengine.String("value"),
	}
	createNodePoolInput := &vke.CreateNodePoolInput{
		AutoScaling:      reqAutoScaling,
		ClientToken:      volcengine.String("BC028527-33B9-4990-A633-84E9F9******"),
		ClusterId:        volcengine.String("cc5silumrsfeq****"),
		KubernetesConfig: reqKubernetesConfig,
		Name:             volcengine.String("test-nodepool"),
		NodeConfig:       reqNodeConfig,
		Tags:             []*vke.TagForCreateNodePoolInput{reqTags1},
	}

	resp, err := svc.CreateNodePool(createNodePoolInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
