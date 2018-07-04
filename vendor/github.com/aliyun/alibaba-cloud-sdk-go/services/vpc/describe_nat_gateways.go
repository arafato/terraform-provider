package vpc

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

// DescribeNatGateways invokes the vpc.DescribeNatGateways API synchronously
// api document: https://help.aliyun.com/api/vpc/describenatgateways.html
func (client *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
	response = CreateDescribeNatGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNatGatewaysWithChan invokes the vpc.DescribeNatGateways API asynchronously
// api document: https://help.aliyun.com/api/vpc/describenatgateways.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNatGatewaysWithChan(request *DescribeNatGatewaysRequest) (<-chan *DescribeNatGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeNatGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNatGateways(request)
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

// DescribeNatGatewaysWithCallback invokes the vpc.DescribeNatGateways API asynchronously
// api document: https://help.aliyun.com/api/vpc/describenatgateways.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNatGatewaysWithCallback(request *DescribeNatGatewaysRequest, callback func(response *DescribeNatGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNatGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeNatGateways(request)
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

// DescribeNatGatewaysRequest is the request struct for api DescribeNatGateways
type DescribeNatGatewaysRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	Name                 string           `position:"Query" name:"Name"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeNatGatewaysResponse is the response struct for api DescribeNatGateways
type DescribeNatGatewaysResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	NatGateways NatGateways `json:"NatGateways" xml:"NatGateways"`
}

// CreateDescribeNatGatewaysRequest creates a request to invoke DescribeNatGateways API
func CreateDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
	request = &DescribeNatGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNatGateways", "vpc", "openAPI")
	return
}

// CreateDescribeNatGatewaysResponse creates a response to parse from DescribeNatGateways response
func CreateDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
	response = &DescribeNatGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
