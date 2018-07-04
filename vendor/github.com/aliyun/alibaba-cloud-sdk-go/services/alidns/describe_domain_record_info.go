package alidns

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

// DescribeDomainRecordInfo invokes the alidns.DescribeDomainRecordInfo API synchronously
// api document: https://help.aliyun.com/api/alidns/describedomainrecordinfo.html
func (client *Client) DescribeDomainRecordInfo(request *DescribeDomainRecordInfoRequest) (response *DescribeDomainRecordInfoResponse, err error) {
	response = CreateDescribeDomainRecordInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRecordInfoWithChan invokes the alidns.DescribeDomainRecordInfo API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedomainrecordinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRecordInfoWithChan(request *DescribeDomainRecordInfoRequest) (<-chan *DescribeDomainRecordInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRecordInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRecordInfo(request)
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

// DescribeDomainRecordInfoWithCallback invokes the alidns.DescribeDomainRecordInfo API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedomainrecordinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRecordInfoWithCallback(request *DescribeDomainRecordInfoRequest, callback func(response *DescribeDomainRecordInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRecordInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRecordInfo(request)
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

// DescribeDomainRecordInfoRequest is the request struct for api DescribeDomainRecordInfo
type DescribeDomainRecordInfoRequest struct {
	*requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
}

// DescribeDomainRecordInfoResponse is the response struct for api DescribeDomainRecordInfo
type DescribeDomainRecordInfoResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainId   string `json:"DomainId" xml:"DomainId"`
	DomainName string `json:"DomainName" xml:"DomainName"`
	PunyCode   string `json:"PunyCode" xml:"PunyCode"`
	GroupId    string `json:"GroupId" xml:"GroupId"`
	GroupName  string `json:"GroupName" xml:"GroupName"`
	RecordId   string `json:"RecordId" xml:"RecordId"`
	RR         string `json:"RR" xml:"RR"`
	Type       string `json:"Type" xml:"Type"`
	Value      string `json:"Value" xml:"Value"`
	TTL        int    `json:"TTL" xml:"TTL"`
	Priority   int    `json:"Priority" xml:"Priority"`
	Line       string `json:"Line" xml:"Line"`
	Status     string `json:"Status" xml:"Status"`
	Locked     bool   `json:"Locked" xml:"Locked"`
}

// CreateDescribeDomainRecordInfoRequest creates a request to invoke DescribeDomainRecordInfo API
func CreateDescribeDomainRecordInfoRequest() (request *DescribeDomainRecordInfoRequest) {
	request = &DescribeDomainRecordInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainRecordInfo", "", "")
	return
}

// CreateDescribeDomainRecordInfoResponse creates a response to parse from DescribeDomainRecordInfo response
func CreateDescribeDomainRecordInfoResponse() (response *DescribeDomainRecordInfoResponse) {
	response = &DescribeDomainRecordInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
