package dcdn

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

// DescribeDcdnService invokes the dcdn.DescribeDcdnService API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnservice.html
func (client *Client) DescribeDcdnService(request *DescribeDcdnServiceRequest) (response *DescribeDcdnServiceResponse, err error) {
	response = CreateDescribeDcdnServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnServiceWithChan invokes the dcdn.DescribeDcdnService API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnServiceWithChan(request *DescribeDcdnServiceRequest) (<-chan *DescribeDcdnServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnService(request)
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

// DescribeDcdnServiceWithCallback invokes the dcdn.DescribeDcdnService API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnServiceWithCallback(request *DescribeDcdnServiceRequest, callback func(response *DescribeDcdnServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnService(request)
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

// DescribeDcdnServiceRequest is the request struct for api DescribeDcdnService
type DescribeDcdnServiceRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnServiceResponse is the response struct for api DescribeDcdnService
type DescribeDcdnServiceResponse struct {
	*responses.BaseResponse
	RequestId          string         `json:"RequestId" xml:"RequestId"`
	InstanceId         string         `json:"InstanceId" xml:"InstanceId"`
	InternetChargeType string         `json:"InternetChargeType" xml:"InternetChargeType"`
	OpeningTime        string         `json:"OpeningTime" xml:"OpeningTime"`
	ChangingChargeType string         `json:"ChangingChargeType" xml:"ChangingChargeType"`
	ChangingAffectTime string         `json:"ChangingAffectTime" xml:"ChangingAffectTime"`
	OperationLocks     OperationLocks `json:"OperationLocks" xml:"OperationLocks"`
}

// CreateDescribeDcdnServiceRequest creates a request to invoke DescribeDcdnService API
func CreateDescribeDcdnServiceRequest() (request *DescribeDcdnServiceRequest) {
	request = &DescribeDcdnServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnService", "dcdn", "openAPI")
	return
}

// CreateDescribeDcdnServiceResponse creates a response to parse from DescribeDcdnService response
func CreateDescribeDcdnServiceResponse() (response *DescribeDcdnServiceResponse) {
	response = &DescribeDcdnServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
