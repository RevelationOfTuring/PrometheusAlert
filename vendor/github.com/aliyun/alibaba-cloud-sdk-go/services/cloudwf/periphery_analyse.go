package cloudwf

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

// PeripheryAnalyse invokes the cloudwf.PeripheryAnalyse API synchronously
// api document: https://help.aliyun.com/api/cloudwf/peripheryanalyse.html
func (client *Client) PeripheryAnalyse(request *PeripheryAnalyseRequest) (response *PeripheryAnalyseResponse, err error) {
	response = CreatePeripheryAnalyseResponse()
	err = client.DoAction(request, response)
	return
}

// PeripheryAnalyseWithChan invokes the cloudwf.PeripheryAnalyse API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/peripheryanalyse.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PeripheryAnalyseWithChan(request *PeripheryAnalyseRequest) (<-chan *PeripheryAnalyseResponse, <-chan error) {
	responseChan := make(chan *PeripheryAnalyseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PeripheryAnalyse(request)
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

// PeripheryAnalyseWithCallback invokes the cloudwf.PeripheryAnalyse API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/peripheryanalyse.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PeripheryAnalyseWithCallback(request *PeripheryAnalyseRequest, callback func(response *PeripheryAnalyseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PeripheryAnalyseResponse
		var err error
		defer close(result)
		response, err = client.PeripheryAnalyse(request)
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

// PeripheryAnalyseRequest is the request struct for api PeripheryAnalyse
type PeripheryAnalyseRequest struct {
	*requests.RpcRequest
	Gsid requests.Integer `position:"Query" name:"Gsid"`
}

// PeripheryAnalyseResponse is the response struct for api PeripheryAnalyse
type PeripheryAnalyseResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreatePeripheryAnalyseRequest creates a request to invoke PeripheryAnalyse API
func CreatePeripheryAnalyseRequest() (request *PeripheryAnalyseRequest) {
	request = &PeripheryAnalyseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "PeripheryAnalyse", "cloudwf", "openAPI")
	return
}

// CreatePeripheryAnalyseResponse creates a response to parse from PeripheryAnalyse response
func CreatePeripheryAnalyseResponse() (response *PeripheryAnalyseResponse) {
	response = &PeripheryAnalyseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}