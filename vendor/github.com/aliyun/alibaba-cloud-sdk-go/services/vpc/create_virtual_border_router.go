package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) CreateVirtualBorderRouter(request *CreateVirtualBorderRouterRequest) (response *CreateVirtualBorderRouterResponse, err error) {
	response = CreateCreateVirtualBorderRouterResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateVirtualBorderRouterWithChan(request *CreateVirtualBorderRouterRequest) (<-chan *CreateVirtualBorderRouterResponse, <-chan error) {
	responseChan := make(chan *CreateVirtualBorderRouterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVirtualBorderRouter(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) CreateVirtualBorderRouterWithCallback(request *CreateVirtualBorderRouterRequest, callback func(response *CreateVirtualBorderRouterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVirtualBorderRouterResponse
		var err error
		defer close(result)
		response, err = client.CreateVirtualBorderRouter(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type CreateVirtualBorderRouterRequest struct {
	*requests.RpcRequest
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
	VbrOwnerId           requests.Integer `position:"Query" name:"VbrOwnerId"`
	VlanId               requests.Integer `position:"Query" name:"VlanId"`
	CircuitCode          string           `position:"Query" name:"CircuitCode"`
	LocalGatewayIp       string           `position:"Query" name:"LocalGatewayIp"`
	PeerGatewayIp        string           `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string           `position:"Query" name:"PeeringSubnetMask"`
	Description          string           `position:"Query" name:"Description"`
	Name                 string           `position:"Query" name:"Name"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type CreateVirtualBorderRouterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	VbrId     string `json:"VbrId" xml:"VbrId"`
}

func CreateCreateVirtualBorderRouterRequest() (request *CreateVirtualBorderRouterRequest) {
	request = &CreateVirtualBorderRouterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVirtualBorderRouter", "vpc", "openAPI")
	return
}

func CreateCreateVirtualBorderRouterResponse() (response *CreateVirtualBorderRouterResponse) {
	response = &CreateVirtualBorderRouterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
