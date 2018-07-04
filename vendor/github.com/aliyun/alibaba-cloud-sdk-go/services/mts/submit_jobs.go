package mts

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

// SubmitJobs invokes the mts.SubmitJobs API synchronously
// api document: https://help.aliyun.com/api/mts/submitjobs.html
func (client *Client) SubmitJobs(request *SubmitJobsRequest) (response *SubmitJobsResponse, err error) {
	response = CreateSubmitJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitJobsWithChan invokes the mts.SubmitJobs API asynchronously
// api document: https://help.aliyun.com/api/mts/submitjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitJobsWithChan(request *SubmitJobsRequest) (<-chan *SubmitJobsResponse, <-chan error) {
	responseChan := make(chan *SubmitJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitJobs(request)
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

// SubmitJobsWithCallback invokes the mts.SubmitJobs API asynchronously
// api document: https://help.aliyun.com/api/mts/submitjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitJobsWithCallback(request *SubmitJobsRequest, callback func(response *SubmitJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitJobsResponse
		var err error
		defer close(result)
		response, err = client.SubmitJobs(request)
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

// SubmitJobsRequest is the request struct for api SubmitJobs
type SubmitJobsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Input                string           `position:"Query" name:"Input"`
	Outputs              string           `position:"Query" name:"Outputs"`
	OutputBucket         string           `position:"Query" name:"OutputBucket"`
	OutputLocation       string           `position:"Query" name:"OutputLocation"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// SubmitJobsResponse is the response struct for api SubmitJobs
type SubmitJobsResponse struct {
	*responses.BaseResponse
	RequestId     string                    `json:"RequestId" xml:"RequestId"`
	JobResultList JobResultListInSubmitJobs `json:"JobResultList" xml:"JobResultList"`
}

// CreateSubmitJobsRequest creates a request to invoke SubmitJobs API
func CreateSubmitJobsRequest() (request *SubmitJobsRequest) {
	request = &SubmitJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitJobs", "mts", "openAPI")
	return
}

// CreateSubmitJobsResponse creates a response to parse from SubmitJobs response
func CreateSubmitJobsResponse() (response *SubmitJobsResponse) {
	response = &SubmitJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
