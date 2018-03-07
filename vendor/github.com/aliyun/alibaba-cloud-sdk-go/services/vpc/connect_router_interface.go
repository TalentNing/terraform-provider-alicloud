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

func (client *Client) ConnectRouterInterface(request *ConnectRouterInterfaceRequest) (response *ConnectRouterInterfaceResponse, err error) {
	response = CreateConnectRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ConnectRouterInterfaceWithChan(request *ConnectRouterInterfaceRequest) (<-chan *ConnectRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *ConnectRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConnectRouterInterface(request)
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

func (client *Client) ConnectRouterInterfaceWithCallback(request *ConnectRouterInterfaceRequest, callback func(response *ConnectRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConnectRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.ConnectRouterInterface(request)
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

type ConnectRouterInterfaceRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
}

type ConnectRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateConnectRouterInterfaceRequest() (request *ConnectRouterInterfaceRequest) {
	request = &ConnectRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ConnectRouterInterface", "vpc", "openAPI")
	return
}

func CreateConnectRouterInterfaceResponse() (response *ConnectRouterInterfaceResponse) {
	response = &ConnectRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
