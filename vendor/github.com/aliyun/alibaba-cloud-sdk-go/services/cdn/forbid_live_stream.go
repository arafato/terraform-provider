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

// ForbidLiveStream invokes the cdn.ForbidLiveStream API synchronously
// api document: https://help.aliyun.com/api/cdn/forbidlivestream.html
func (client *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
	response = CreateForbidLiveStreamResponse()
	err = client.DoAction(request, response)
	return
}

// ForbidLiveStreamWithChan invokes the cdn.ForbidLiveStream API asynchronously
// api document: https://help.aliyun.com/api/cdn/forbidlivestream.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ForbidLiveStreamWithChan(request *ForbidLiveStreamRequest) (<-chan *ForbidLiveStreamResponse, <-chan error) {
	responseChan := make(chan *ForbidLiveStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ForbidLiveStream(request)
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

// ForbidLiveStreamWithCallback invokes the cdn.ForbidLiveStream API asynchronously
// api document: https://help.aliyun.com/api/cdn/forbidlivestream.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ForbidLiveStreamWithCallback(request *ForbidLiveStreamRequest, callback func(response *ForbidLiveStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ForbidLiveStreamResponse
		var err error
		defer close(result)
		response, err = client.ForbidLiveStream(request)
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

// ForbidLiveStreamRequest is the request struct for api ForbidLiveStream
type ForbidLiveStreamRequest struct {
	*requests.RpcRequest
	ResumeTime     string           `position:"Query" name:"ResumeTime"`
	AppName        string           `position:"Query" name:"AppName"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	LiveStreamType string           `position:"Query" name:"LiveStreamType"`
	DomainName     string           `position:"Query" name:"DomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	StreamName     string           `position:"Query" name:"StreamName"`
}

// ForbidLiveStreamResponse is the response struct for api ForbidLiveStream
type ForbidLiveStreamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateForbidLiveStreamRequest creates a request to invoke ForbidLiveStream API
func CreateForbidLiveStreamRequest() (request *ForbidLiveStreamRequest) {
	request = &ForbidLiveStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ForbidLiveStream", "", "")
	return
}

// CreateForbidLiveStreamResponse creates a response to parse from ForbidLiveStream response
func CreateForbidLiveStreamResponse() (response *ForbidLiveStreamResponse) {
	response = &ForbidLiveStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
