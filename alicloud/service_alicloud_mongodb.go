package alicloud

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type MongoDBInstance struct {
	ChargeType            string `json:"ChargeType"`
	CreationTime          string `json:"CreationTime"`
	DBInstanceClass       string `json:"DBInstanceClass"`
	DBInstanceDescription string `json:"DBInstanceDescription"`
	DBInstanceID          string `json:"DBInstanceId"`
	DBInstanceStatus      string `json:"DBInstanceStatus"`
	DBInstanceStorage     int    `json:"DBInstanceStorage"`
	DBInstanceType        string `json:"DBInstanceType"`
	Engine                string `json:"Engine"`
	EngineVersion         string `json:"EngineVersion"`
	ExpireTime            string `json:"ExpireTime"`
	LockMode              string `json:"LockMode"`
	NetworkType           string `json:"NetworkType"`
	RegionID              string `json:"RegionId"`
	ReplicationFactor     string `json:"ReplicationFactor"`
	ZoneID                string `json:"ZoneId"`
}

type ItemsInDescribeMongoDBInstances struct {
	DBInstances []MongoDBInstance `json:"DBInstance"`
}

type DescribeMongoDBInstancesResponse struct {
	PageNumber int                             `json:"PageNumber"`
	PageSize   int                             `json:"PageSize"`
	RequestID  string                          `json:"RequestId"`
	TotalCount int                             `json:"TotalCount"`
	Items      ItemsInDescribeMongoDBInstances `json:"DBInstances"`
}

func (client *AliyunClient) DescribeMongoDBInstances(request *requests.CommonRequest) (response *DescribeMongoDBInstancesResponse, err error) {
	request.Version = ApiVersion20151201
	request.ApiName = "DescribeDBInstances"
	resp, err := client.ecsconn.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}
	response = new(DescribeMongoDBInstancesResponse)
	err = json.Unmarshal(resp.BaseResponse.GetHttpContentBytes(), &response)

	return response, err
}
