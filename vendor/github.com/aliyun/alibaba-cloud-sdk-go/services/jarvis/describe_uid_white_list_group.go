package jarvis

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

// DescribeUidWhiteListGroup invokes the jarvis.DescribeUidWhiteListGroup API synchronously
// api document: https://help.aliyun.com/api/jarvis/describeuidwhitelistgroup.html
func (client *Client) DescribeUidWhiteListGroup(request *DescribeUidWhiteListGroupRequest) (response *DescribeUidWhiteListGroupResponse, err error) {
	response = CreateDescribeUidWhiteListGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUidWhiteListGroupWithChan invokes the jarvis.DescribeUidWhiteListGroup API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describeuidwhitelistgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUidWhiteListGroupWithChan(request *DescribeUidWhiteListGroupRequest) (<-chan *DescribeUidWhiteListGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeUidWhiteListGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUidWhiteListGroup(request)
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

// DescribeUidWhiteListGroupWithCallback invokes the jarvis.DescribeUidWhiteListGroup API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describeuidwhitelistgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUidWhiteListGroupWithCallback(request *DescribeUidWhiteListGroupRequest, callback func(response *DescribeUidWhiteListGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUidWhiteListGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeUidWhiteListGroup(request)
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

// DescribeUidWhiteListGroupRequest is the request struct for api DescribeUidWhiteListGroup
type DescribeUidWhiteListGroupRequest struct {
	*requests.RpcRequest
	SourceIp      string           `position:"Query" name:"SourceIp"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	WhiteListType requests.Integer `position:"Query" name:"WhiteListType"`
	DstIP         string           `position:"Query" name:"DstIP"`
	Lang          string           `position:"Query" name:"Lang"`
	SrcUid        string           `position:"Query" name:"SrcUid"`
	Status        string           `position:"Query" name:"Status"`
	SourceCode    string           `position:"Query" name:"SourceCode"`
}

// DescribeUidWhiteListGroupResponse is the response struct for api DescribeUidWhiteListGroup
type DescribeUidWhiteListGroupResponse struct {
	*responses.BaseResponse
	RequestId   string   `json:"RequestId" xml:"RequestId"`
	Module      string   `json:"module" xml:"module"`
	ProductList []string `json:"ProductList" xml:"ProductList"`
	PageInfo    PageInfo `json:"PageInfo" xml:"PageInfo"`
	DataList    []Data   `json:"DataList" xml:"DataList"`
}

// CreateDescribeUidWhiteListGroupRequest creates a request to invoke DescribeUidWhiteListGroup API
func CreateDescribeUidWhiteListGroupRequest() (request *DescribeUidWhiteListGroupRequest) {
	request = &DescribeUidWhiteListGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeUidWhiteListGroup", "", "")
	return
}

// CreateDescribeUidWhiteListGroupResponse creates a response to parse from DescribeUidWhiteListGroup response
func CreateDescribeUidWhiteListGroupResponse() (response *DescribeUidWhiteListGroupResponse) {
	response = &DescribeUidWhiteListGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
