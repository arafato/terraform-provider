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

// DescribeDomainHttpsData invokes the cdn.DescribeDomainHttpsData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainhttpsdata.html
func (client *Client) DescribeDomainHttpsData(request *DescribeDomainHttpsDataRequest) (response *DescribeDomainHttpsDataResponse, err error) {
	response = CreateDescribeDomainHttpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainHttpsDataWithChan invokes the cdn.DescribeDomainHttpsData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainhttpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainHttpsDataWithChan(request *DescribeDomainHttpsDataRequest) (<-chan *DescribeDomainHttpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainHttpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainHttpsData(request)
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

// DescribeDomainHttpsDataWithCallback invokes the cdn.DescribeDomainHttpsData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainhttpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainHttpsDataWithCallback(request *DescribeDomainHttpsDataRequest, callback func(response *DescribeDomainHttpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainHttpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainHttpsData(request)
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

// DescribeDomainHttpsDataRequest is the request struct for api DescribeDomainHttpsData
type DescribeDomainHttpsDataRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
	DomainType    string           `position:"Query" name:"DomainType"`
	TimeMerge     string           `position:"Query" name:"TimeMerge"`
	Interval      string           `position:"Query" name:"Interval"`
	Cls           string           `position:"Query" name:"Cls"`
	FixTimeGap    string           `position:"Query" name:"FixTimeGap"`
}

// DescribeDomainHttpsDataResponse is the response struct for api DescribeDomainHttpsData
type DescribeDomainHttpsDataResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	DomainNames          string               `json:"DomainNames" xml:"DomainNames"`
	DataInterval         string               `json:"DataInterval" xml:"DataInterval"`
	HttpsStatisticsInfos HttpsStatisticsInfos `json:"HttpsStatisticsInfos" xml:"HttpsStatisticsInfos"`
}

// CreateDescribeDomainHttpsDataRequest creates a request to invoke DescribeDomainHttpsData API
func CreateDescribeDomainHttpsDataRequest() (request *DescribeDomainHttpsDataRequest) {
	request = &DescribeDomainHttpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainHttpsData", "", "")
	return
}

// CreateDescribeDomainHttpsDataResponse creates a response to parse from DescribeDomainHttpsData response
func CreateDescribeDomainHttpsDataResponse() (response *DescribeDomainHttpsDataResponse) {
	response = &DescribeDomainHttpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
