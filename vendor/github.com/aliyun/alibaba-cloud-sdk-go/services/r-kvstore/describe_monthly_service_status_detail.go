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

// DescribeMonthlyServiceStatusDetail invokes the r_kvstore.DescribeMonthlyServiceStatusDetail API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describemonthlyservicestatusdetail.html
func (client *Client) DescribeMonthlyServiceStatusDetail(request *DescribeMonthlyServiceStatusDetailRequest) (response *DescribeMonthlyServiceStatusDetailResponse, err error) {
	response = CreateDescribeMonthlyServiceStatusDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonthlyServiceStatusDetailWithChan invokes the r_kvstore.DescribeMonthlyServiceStatusDetail API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describemonthlyservicestatusdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonthlyServiceStatusDetailWithChan(request *DescribeMonthlyServiceStatusDetailRequest) (<-chan *DescribeMonthlyServiceStatusDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeMonthlyServiceStatusDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonthlyServiceStatusDetail(request)
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

// DescribeMonthlyServiceStatusDetailWithCallback invokes the r_kvstore.DescribeMonthlyServiceStatusDetail API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describemonthlyservicestatusdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonthlyServiceStatusDetailWithCallback(request *DescribeMonthlyServiceStatusDetailRequest, callback func(response *DescribeMonthlyServiceStatusDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonthlyServiceStatusDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonthlyServiceStatusDetail(request)
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

// DescribeMonthlyServiceStatusDetailRequest is the request struct for api DescribeMonthlyServiceStatusDetail
type DescribeMonthlyServiceStatusDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Month                string           `position:"Query" name:"Month"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeMonthlyServiceStatusDetailResponse is the response struct for api DescribeMonthlyServiceStatusDetail
type DescribeMonthlyServiceStatusDetailResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	InstanceId    string        `json:"InstanceId" xml:"InstanceId"`
	UptimePct     float64       `json:"UptimePct" xml:"UptimePct"`
	AffectedInfos AffectedInfos `json:"AffectedInfos" xml:"AffectedInfos"`
}

// CreateDescribeMonthlyServiceStatusDetailRequest creates a request to invoke DescribeMonthlyServiceStatusDetail API
func CreateDescribeMonthlyServiceStatusDetailRequest() (request *DescribeMonthlyServiceStatusDetailRequest) {
	request = &DescribeMonthlyServiceStatusDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeMonthlyServiceStatusDetail", "redisa", "openAPI")
	return
}

// CreateDescribeMonthlyServiceStatusDetailResponse creates a response to parse from DescribeMonthlyServiceStatusDetail response
func CreateDescribeMonthlyServiceStatusDetailResponse() (response *DescribeMonthlyServiceStatusDetailResponse) {
	response = &DescribeMonthlyServiceStatusDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
