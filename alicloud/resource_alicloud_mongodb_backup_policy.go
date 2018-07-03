package alicloud

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlicloudMongoDBBackupPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudMongoDBBackupPolicyCreate,
		Read:   resourceAlicloudMongoDBBackupPolicyRead,
		Update: resourceAlicloudMongoDBBackupPolicyUpdate,
		Delete: resourceAlicloudMongoDBBackupPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"preferred_backup_time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"preferred_backup_period": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateAllowedStringValue([]string{
					"Monday",
					"Tuesday",
					"Wednesday",
					"Thursday",
					"Friday",
					"Saturday",
					"Sunday",
				}),
			},
		},
	}
}

func resourceAlicloudMongoDBBackupPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)

	request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
	request.RegionId = getRegionId(d, meta)
	request.QueryParams["DBInstanceId"] = d.Get("instance_id").(string)
	request.QueryParams["PreferredBackupTime"] = d.Get("preferred_backup_time").(string)
	request.QueryParams["PreferredBackupPeriod"] = d.Get("preferred_backup_period").(string)

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		if err := client.ModifyMongoDBBackupPolicy(request); err != nil {
			return resource.NonRetryableError(fmt.Errorf("Create security whitelist ips got an error: %#v", err))
		}
		return nil
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s%s%s", request.QueryParams["DBInstanceId"], COLON_SEPARATED, resource.UniqueId()))
	return resourceAlicloudMongoDBBackupPolicyRead(d, meta)
}

func resourceAlicloudMongoDBBackupPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	instanceID := strings.Split(d.Id(), COLON_SEPARATED)[0]

	request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
	request.RegionId = getRegionId(d, meta)
	request.QueryParams["DBInstanceId"] = instanceID
	policy, err := client.DescribeMongoDBBackupPolicy(request)
	if err != nil {
		if NotFoundDBInstance(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Describe MongoDB Backup Policy: %#v", err)
	}
	if policy == nil {
		d.SetId("")
		return nil
	}

	d.Set("instance_id", instanceID)
	d.Set("preferred_backup_time", policy.PreferredBackupTime)
	d.Set("preferred_backup_period", policy.PreferredBackupPeriod)

	return nil
}

func resourceAlicloudMongoDBBackupPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	instanceID := strings.Split(d.Id(), COLON_SEPARATED)[0]

	if d.HasChange("preferred_backup_time") && d.HasChange("preferred_backup_period") {
		request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
		request.RegionId = getRegionId(d, meta)
		request.QueryParams["DBInstanceId"] = instanceID
		request.QueryParams["PreferredBackupTime"] = d.Get("preferred_backup_time").(string)
		request.QueryParams["PreferredBackupPeriod"] = d.Get("preferred_backup_period").(string)
		if err := client.ModifyMongoDBBackupPolicy(request); err != nil {
			return err
		}
	}

	return resourceAlicloudMongoDBBackupPolicyRead(d, meta)
}

func resourceAlicloudMongoDBBackupPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	// There is no explicit delete, only update with modified security ips
	return nil
}
