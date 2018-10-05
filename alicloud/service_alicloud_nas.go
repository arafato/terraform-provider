package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
)

func (client *AliyunClient) DescribeNASFilesystemById(id string) (instance *nas.FileSystem, err error) {
	conn := client.nasconn
	request := nas.CreateDescribeFileSystemsRequest()
	request.FileSystemId = id

	resp, err := conn.DescribeFileSystems(request)
	if err != nil {
		if IsExceptedError(err, InvalidNASFileSystem) {
			return nil, GetNotFoundErrorFromString(GetNotFoundMessage("NAS Filesystem", id))
		}
		return nil, err
	}

	if resp == nil || len(resp.FileSystems.FileSystem) <= 0 {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("NAS Filesystem", id))
	}

	return &resp.FileSystems.FileSystem[0], nil
}
