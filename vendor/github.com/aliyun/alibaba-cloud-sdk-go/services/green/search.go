package green

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

// Search invokes the green.Search API synchronously
// api document: https://help.aliyun.com/api/green/search.html
func (client *Client) Search(request *SearchRequest) (response *SearchResponse, err error) {
	response = CreateSearchResponse()
	err = client.DoAction(request, response)
	return
}

// SearchWithChan invokes the green.Search API asynchronously
// api document: https://help.aliyun.com/api/green/search.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWithChan(request *SearchRequest) (<-chan *SearchResponse, <-chan error) {
	responseChan := make(chan *SearchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Search(request)
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

// SearchWithCallback invokes the green.Search API asynchronously
// api document: https://help.aliyun.com/api/green/search.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWithCallback(request *SearchRequest, callback func(response *SearchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchResponse
		var err error
		defer close(result)
		response, err = client.Search(request)
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

// SearchRequest is the request struct for api Search
type SearchRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// SearchResponse is the response struct for api Search
type SearchResponse struct {
	*responses.BaseResponse
}

// CreateSearchRequest creates a request to invoke Search API
func CreateSearchRequest() (request *SearchRequest) {
	request = &SearchRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-25", "Search", "/green/sface/search", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchResponse creates a response to parse from Search response
func CreateSearchResponse() (response *SearchResponse) {
	response = &SearchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
