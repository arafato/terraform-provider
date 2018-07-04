package r_kvstore

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

// DescribeInstanceConfig invokes the r_kvstore.DescribeInstanceConfig API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstanceconfig.html
func (client *Client) DescribeInstanceConfig(request *DescribeInstanceConfigRequest) (response *DescribeInstanceConfigResponse, err error) {
	response = CreateDescribeInstanceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceConfigWithChan invokes the r_kvstore.DescribeInstanceConfig API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstanceconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceConfigWithChan(request *DescribeInstanceConfigRequest) (<-chan *DescribeInstanceConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceConfig(request)
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

// DescribeInstanceConfigWithCallback invokes the r_kvstore.DescribeInstanceConfig API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstanceconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceConfigWithCallback(request *DescribeInstanceConfigRequest, callback func(response *DescribeInstanceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceConfig(request)
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

// DescribeInstanceConfigRequest is the request struct for api DescribeInstanceConfig
type DescribeInstanceConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeInstanceConfigResponse is the response struct for api DescribeInstanceConfig
type DescribeInstanceConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Config    string `json:"Config" xml:"Config"`
}

// CreateDescribeInstanceConfigRequest creates a request to invoke DescribeInstanceConfig API
func CreateDescribeInstanceConfigRequest() (request *DescribeInstanceConfigRequest) {
	request = &DescribeInstanceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceConfig", "redisa", "openAPI")
	return
}

// CreateDescribeInstanceConfigResponse creates a response to parse from DescribeInstanceConfig response
func CreateDescribeInstanceConfigResponse() (response *DescribeInstanceConfigResponse) {
	response = &DescribeInstanceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
