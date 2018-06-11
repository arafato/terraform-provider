package cdn

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

// DescribeLiveAppRecordConfig invokes the cdn.DescribeLiveAppRecordConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/describeliveapprecordconfig.html
func (client *Client) DescribeLiveAppRecordConfig(request *DescribeLiveAppRecordConfigRequest) (response *DescribeLiveAppRecordConfigResponse, err error) {
	response = CreateDescribeLiveAppRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveAppRecordConfigWithChan invokes the cdn.DescribeLiveAppRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeliveapprecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveAppRecordConfigWithChan(request *DescribeLiveAppRecordConfigRequest) (<-chan *DescribeLiveAppRecordConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveAppRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveAppRecordConfig(request)
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

// DescribeLiveAppRecordConfigWithCallback invokes the cdn.DescribeLiveAppRecordConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/describeliveapprecordconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveAppRecordConfigWithCallback(request *DescribeLiveAppRecordConfigRequest, callback func(response *DescribeLiveAppRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveAppRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveAppRecordConfig(request)
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

// DescribeLiveAppRecordConfigRequest is the request struct for api DescribeLiveAppRecordConfig
type DescribeLiveAppRecordConfigRequest struct {
	*requests.RpcRequest
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveAppRecordConfigResponse is the response struct for api DescribeLiveAppRecordConfig
type DescribeLiveAppRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	LiveAppRecord LiveAppRecord `json:"LiveAppRecord" xml:"LiveAppRecord"`
}

// CreateDescribeLiveAppRecordConfigRequest creates a request to invoke DescribeLiveAppRecordConfig API
func CreateDescribeLiveAppRecordConfigRequest() (request *DescribeLiveAppRecordConfigRequest) {
	request = &DescribeLiveAppRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveAppRecordConfig", "", "")
	return
}

// CreateDescribeLiveAppRecordConfigResponse creates a response to parse from DescribeLiveAppRecordConfig response
func CreateDescribeLiveAppRecordConfigResponse() (response *DescribeLiveAppRecordConfigResponse) {
	response = &DescribeLiveAppRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
