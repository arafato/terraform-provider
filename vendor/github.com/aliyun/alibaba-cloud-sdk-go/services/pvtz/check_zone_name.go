package pvtz

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

// CheckZoneName invokes the pvtz.CheckZoneName API synchronously
// api document: https://help.aliyun.com/api/pvtz/checkzonename.html
func (client *Client) CheckZoneName(request *CheckZoneNameRequest) (response *CheckZoneNameResponse, err error) {
	response = CreateCheckZoneNameResponse()
	err = client.DoAction(request, response)
	return
}

// CheckZoneNameWithChan invokes the pvtz.CheckZoneName API asynchronously
// api document: https://help.aliyun.com/api/pvtz/checkzonename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckZoneNameWithChan(request *CheckZoneNameRequest) (<-chan *CheckZoneNameResponse, <-chan error) {
	responseChan := make(chan *CheckZoneNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckZoneName(request)
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

// CheckZoneNameWithCallback invokes the pvtz.CheckZoneName API asynchronously
// api document: https://help.aliyun.com/api/pvtz/checkzonename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckZoneNameWithCallback(request *CheckZoneNameRequest, callback func(response *CheckZoneNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckZoneNameResponse
		var err error
		defer close(result)
		response, err = client.CheckZoneName(request)
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

// CheckZoneNameRequest is the request struct for api CheckZoneName
type CheckZoneNameRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	ZoneName     string `position:"Query" name:"ZoneName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
}

// CheckZoneNameResponse is the response struct for api CheckZoneName
type CheckZoneNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Check     bool   `json:"Check" xml:"Check"`
}

// CreateCheckZoneNameRequest creates a request to invoke CheckZoneName API
func CreateCheckZoneNameRequest() (request *CheckZoneNameRequest) {
	request = &CheckZoneNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "CheckZoneName", "pvtz", "openAPI")
	return
}

// CreateCheckZoneNameResponse creates a response to parse from CheckZoneName response
func CreateCheckZoneNameResponse() (response *CheckZoneNameResponse) {
	response = &CheckZoneNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
