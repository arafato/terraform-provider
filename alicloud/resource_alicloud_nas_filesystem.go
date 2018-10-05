package alicloud

import (
	"fmt"
	"strings"
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

	conn.DescribeFileSystems()

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
	conn := client.rkvconn
	d.Partial(true)

	if d.HasChange("security_ips") {
		request := r_kvstore.CreateModifySecurityIpsRequest()
		request.SecurityIpGroupName = "default"
		request.InstanceId = d.Id()
		if len(d.Get("security_ips").(*schema.Set).List()) > 0 {
			request.SecurityIps = strings.Join(expandStringList(d.Get("security_ips").(*schema.Set).List())[:], COMMA_SEPARATED)
		} else {
			return fmt.Errorf("Security ips cannot be empty")
		}
		// wait instance status is Normal before modifying
		if err := client.WaitForRKVInstance(d.Id(), Normal, DefaultLongTimeout); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}
		if _, err := conn.ModifySecurityIps(request); err != nil {
			return fmt.Errorf("Create security whitelist ips got an error: %#v", err)
		}
		d.SetPartial("security_ips")
		// wait instance status is Normal after modifying
		if err := client.WaitForRKVInstance(d.Id(), Normal, DefaultLongTimeout); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}
	}

	if d.IsNewResource() {
		d.Partial(false)
		return resourceAlicloudNASFilesystemRead(d, meta)
	}

	if d.HasChange("instance_class") {
		request := r_kvstore.CreateModifyInstanceSpecRequest()
		request.InstanceId = d.Id()
		request.InstanceClass = d.Get("instance_class").(string)
		request.EffectiveTime = "Immediately"
		if _, err := conn.ModifyInstanceSpec(request); err != nil {
			return err
		}
		// wait instance status is Normal after modifying
		if err := client.WaitForRKVInstance(d.Id(), Normal, DefaultLongTimeout); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}

		d.SetPartial("instance_class")
	}

	request := r_kvstore.CreateModifyInstanceAttributeRequest()
	request.InstanceId = d.Id()
	update := false
	if d.HasChange("instance_name") {
		request.InstanceName = d.Get("instance_name").(string)
		update = true

		d.SetPartial("instance_name")
	}

	if d.HasChange("password") {
		request.NewPassword = d.Get("password").(string)
		update = true
		d.SetPartial("password")
	}

	if update {
		// wait instance status is Normal before modifying
		if err := client.WaitForRKVInstance(d.Id(), Normal, DefaultLongTimeout); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}
		if _, err := conn.ModifyInstanceAttribute(request); err != nil {
			return fmt.Errorf("ModifyRKVInstanceDescription got an error: %#v", err)
		}
		// wait instance status is Normal after modifying
		if err := client.WaitForRKVInstance(d.Id(), Normal, DefaultLongTimeout); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}
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

	instance, err := client.DescribeRKVInstanceById(d.Id())
	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return fmt.Errorf("Error Describe KVStore InstanceAttribute: %#v", err)
	}
	if PayType(instance.ChargeType) == Prepaid {
		return fmt.Errorf("At present, 'Prepaid' instance cannot be deleted and must wait it to be expired and release it automatically")
	}
	request := r_kvstore.CreateDeleteInstanceRequest()
	request.InstanceId = d.Id()

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
