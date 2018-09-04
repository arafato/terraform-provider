package ots

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

// UpdateInstance invokes the ots.UpdateInstance API synchronously
// api document: https://help.aliyun.com/api/ots/updateinstance.html
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
	response = CreateUpdateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateInstanceWithChan invokes the ots.UpdateInstance API asynchronously
// api document: https://help.aliyun.com/api/ots/updateinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateInstanceWithChan(request *UpdateInstanceRequest) (<-chan *UpdateInstanceResponse, <-chan error) {
	responseChan := make(chan *UpdateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateInstance(request)
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

// UpdateInstanceWithCallback invokes the ots.UpdateInstance API asynchronously
// api document: https://help.aliyun.com/api/ots/updateinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateInstanceWithCallback(request *UpdateInstanceRequest, callback func(response *UpdateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateInstanceResponse
		var err error
		defer close(result)
		response, err = client.UpdateInstance(request)
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

// UpdateInstanceRequest is the request struct for api UpdateInstance
type UpdateInstanceRequest struct {
	*requests.RpcRequest
	AccessKeyId     string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceName    string           `position:"Query" name:"InstanceName"`
	Network         string           `position:"Query" name:"Network"`
}

// UpdateInstanceResponse is the response struct for api UpdateInstance
type UpdateInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateInstanceRequest creates a request to invoke UpdateInstance API
func CreateUpdateInstanceRequest() (request *UpdateInstanceRequest) {
	request = &UpdateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ots", "2016-06-20", "UpdateInstance", "ots", "openAPI")
	return
}

// CreateUpdateInstanceResponse creates a response to parse from UpdateInstance response
func CreateUpdateInstanceResponse() (response *UpdateInstanceResponse) {
	response = &UpdateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
