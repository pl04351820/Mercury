package TencentCloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	scf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
)

type ScfClient struct {
	Svc *scf.Client
}

func NewScfClient() ScfClient {
	var newClient ScfClient

	// For now. It does not support for reading from environment variable.
	// TODO: Add this into config file.
	credential := common.NewCredential(
		"xxxxxxxxxxxxxxx",
		"xxxxxxxxxxxxxxx",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.HttpProfile.ReqTimeout = 10
	cpf.SignMethod = "HmacSHA1"

	client, _ := scf.NewClient(credential, "ap-guangzhou", cpf)
	newClient.Svc = client

	return newClient
}

func (s *ScfClient) Invoke(functionName string, input []byte) ([]byte, error) {
	payload := string(input)

	request := scf.NewInvokeRequest()
	request.FunctionName = &functionName
	request.ClientContext = &payload

	response, err := s.Svc.Invoke(request)
	return []byte(*response.Response.Result.RetMsg), err
}
