package alicloud

import (
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlicloudNASFilesystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudNASFilesystemCreate,
		Read:   resourceAlicloudNASFilesystemRead,
		Update: resourceAlicloudNASFilesystemUpdate,
		Delete: resourceAlicloudNASFilesystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"storage_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateAllowedStringValue([]string{
					"Capacity",
					"Performance",
				}),
			},
			"protocol_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateAllowedStringValue([]string{
					"NFS",
					"SMB",
				}),
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateNASDescription,
			},
			"create_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"metered_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudNASFilesystemCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	conn := client.nasconn

	request := nas.CreateCreateFileSystemRequest()
	request.StorageType = d.Get("storage_type").(string)
	request.ProtocolType = d.Get("protocol_type").(string)
	request.Description = d.Get("description").(string)

	resp, err := conn.CreateFileSystem(request)

	if err != nil {
		return fmt.Errorf("Error creating Alicloud NAS Filesytem: %#v", err)
	}

	d.SetId(resp.FileSystemId)

	return resourceAlicloudNASFilesystemRead(d, meta)
}

func resourceAlicloudNASFilesystemUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	conn := client.nasconn
	d.Partial(true)

	if d.HasChange("description") {
		request := nas.CreateModifyFileSystemRequest()
		request.Description = d.Get("description").(string)
		conn.ModifyFileSystem(request)
		d.SetPartial("description")
	}

	d.Partial(false)
	return resourceAlicloudNASFilesystemRead(d, meta)
}

func resourceAlicloudNASFilesystemRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	conn := client.nasconn

	fs, err := client.DescribeNASFilesystemById(d.Id())

	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Describe NAS Filesystem: %#v", err)
	}

	d.Set("create_time", fs.CreateTime)
	d.Set("description", fs.Destription)
	d.Set("metered_size", fs.MeteredSize)
	d.Set("protocol_type", fs.ProtocolType)
	d.Set("storage_type", fs.StorageType)

	return nil
}

func resourceAlicloudNASFilesystemDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)

	fs, err := client.DescribeNASFilesystemById(d.Id())
	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return fmt.Errorf("Error DescribeNASFilesystemById : %#v", err)
	}

	// Check if filesystem has mount points attached

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		_, err := client.rkvconn.DeleteInstance(request)

		if err != nil {
			if IsExceptedError(err, InvalidNASFilesystemIdNotFound) {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("Delete KVStore instance timeout and got an error: %#v", err))
		}

		if _, err := client.DescribeRKVInstanceById(d.Id()); err != nil {
			if NotFoundError(err) {
				return nil
			}
			return resource.NonRetryableError(fmt.Errorf("Error Describe KVStore InstanceAttribute: %#v", err))
		}

		return resource.RetryableError(fmt.Errorf("Delete KVStore instance timeout and got an error: %#v", err))
	})
}
