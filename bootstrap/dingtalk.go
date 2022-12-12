package bootstrap

import (
	"fmt"
	"strings"

	facebody "github.com/alibabacloud-go/facebody-20191230/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	credential "github.com/aliyun/credentials-go/credentials"
)

func StartDingTalk() {
	config := new(rpc.Config)

	// init config with ak
	config.SetAccessKeyId("ACCESS_KEY_ID").
		SetAccessKeySecret("ACCESS_KEY_SECRET").
		SetRegionId("cn-hangzhou").
		SetEndpoint("facebody.cn-hangzhou.aliyuncs.com")

	// init config with credential
	credentialConfig := &credential.Config{
		AccessKeyId:     config.AccessKeyId,
		AccessKeySecret: config.AccessKeySecret,
		SecurityToken:   config.SecurityToken,
	}
	// If you have any questions, please refer to it https://github.com/aliyun/credentials-go/blob/master/README-CN.md
	cred, err := credential.NewCredential(credentialConfig)
	if err != nil {
		panic(err)
	}
	config.SetCredential(cred).
		SetEndpoint("facebody.cn-hangzhou.aliyuncs.com")

	// init client
	client, err := facebody.NewClient(config)
	if err != nil {
		panic(err)
	}

	// init runtimeObject
	runtimeObject := new(util.RuntimeOptions).SetAutoretry(false).
		SetMaxIdleConns(3)

	// init request
	request := new(facebody.DetectFaceRequest)

	// call api
	resp, err := client.DetectFace(request, runtimeObject)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

	// file upload
	uploadRequest := new(facebody.DetectFaceAdvanceRequest).SetImageURLObject(strings.NewReader("demo"))

	// call api
	uploadResp, err := client.DetectFaceAdvance(uploadRequest, runtimeObject)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(uploadResp)

}
