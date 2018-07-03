package alicloud

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlicloudMongoDBInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudMongoDBInstanceCreate,
		Read:   resourceAlicloudMongoDBInstanceRead,
		Update: resourceAlicloudMongoDBInstanceUpdate,
		Delete: resourceAlicloudMongoDBInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"engine_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue([]string{"3.2", "3.4"}),
				ForceNew:     true,
				Required:     true,
			},
			"storage_engine": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue([]string{"WiredTiger", "RocksDB"}),
				Optional:     true,
				Default:      "WiredTiger",
			},
			"instance_class": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_storage": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_ips": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},

			"instance_charge_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue([]string{string(PostPaid), string(PrePaid)}),
				Optional:     true,
				ForceNew:     true,
				Default:      PostPaid,
			},
			"period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36}),
				Optional:     true,
				Default:      1,
			},
			"network_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue([]string{"Classic", "VPC"}),
				Optional:     true,
				ForceNew:     true,
				Default:      "Classic",
			},
			"zone_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"vswitch_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudMongoDBInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)

	request, err := buildMongoDBCreateRequest(d, meta)

	if err != nil {
		return err
	}

	resp, err := client.CreateMongoDBInstance(request)

	if err != nil {
		return fmt.Errorf("Error creating Alicloud MongoDB instance: %#v", err)
	}

	d.SetId(resp.DBInstanceId)

	// wait instance status change from Creating to running
	if err := client.WaitForMongoDBInstance(d.Id(), getRegionId(d, meta), Running, DefaultLongTimeout); err != nil {
		return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
	}

	return resourceAlicloudMongoDBInstanceRead(d, meta)
}

func resourceAlicloudMongoDBInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)
	d.Partial(true)

	if d.HasChange("security_ips") {
		ipList := expandStringList(d.Get("security_ips").(*schema.Set).List())
		ipstr := strings.Join(ipList[:], COMMA_SEPARATED)
		// default disable connect from outside
		if ipstr == "" {
			ipstr = LOCAL_HOST_IP
		}
		request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
		request.RegionId = getRegionId(d, meta)
		request.QueryParams["SecurityIps"] = ipstr
		request.QueryParams["DBInstanceId"] = d.Id()
		if err := client.ModifyMongoDBSecurityIps(request); err != nil {
			return fmt.Errorf("Modify MongoDB security ips %s got an error: %#v", ipstr, err)
		}
		d.SetPartial("security_ips")
	}

	update := false
	request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
	request.RegionId = getRegionId(d, meta)
	request.QueryParams["DBInstanceId"] = d.Id()

	if d.HasChange("instance_storage") {
		request.QueryParams["DBInstanceStorage"] = strconv.Itoa(d.Get("instance_storage").(int))
		update = true
		d.SetPartial("instance_storage")
	}

	if d.HasChange("instance_class") {
		request.QueryParams["DBInstanceClass"] = d.Get("instance_class").(string)
		update = true
		d.SetPartial("instance_class")
	}

	if update {
		// wait instance status is running before modifying
		if err := client.WaitForMongoDBInstance(d.Id(), getRegionId(d, meta), Running, 500); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}
		if err := client.ModifyMongoDBInstanceSpec(request); err != nil {
			return err
		}
		// wait instance status is running after modifying
		if err := client.WaitForMongoDBInstance(d.Id(), getRegionId(d, meta), Running, 500); err != nil {
			return fmt.Errorf("WaitForInstance %s got error: %#v", Running, err)
		}
	}

	if d.HasChange("description") {
		request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
		request.RegionId = getRegionId(d, meta)
		request.QueryParams["DBInstanceId"] = d.Id()
		request.QueryParams["DBInstanceDescription"] = d.Get("description").(string)
		if err := client.ModifyMongoDBInstanceDescription(request); err != nil {
			return fmt.Errorf("ModifyMongoDBInstanceDescription got an error: %#v", err)
		}
	}

	d.Partial(false)
	return resourceAlicloudMongoDBInstanceRead(d, meta)
}

func resourceAlicloudMongoDBInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)

	instance, err := client.DescribeMongoDBInstanceById(d.Id(), getRegionId(d, meta))
	if err != nil {
		if NotFoundDBInstance(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Describe DB InstanceAttribute: %#v", err)
	}

	request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
	request.RegionId = getRegionId(d, meta)
	request.QueryParams["DBInstanceId"] = d.Id()
	ips, err := client.DescribeMongoDBSecurityIps(request)
	if err != nil {
		return fmt.Errorf("[ERROR] Describe DB security ips error: %#v", err)
	}
	d.SetId(instance.DBInstanceID)
	d.Set("security_ips", strings.Split(ips.SecurityIps, COMMA_SEPARATED))
	d.Set("engine_version", instance.EngineVersion)
	d.Set("instance_class", instance.DBInstanceClass)
	d.Set("instance_storage", instance.DBInstanceStorage)
	d.Set("instance_charge_type", instance.ChargeType)
	d.Set("period", d.Get("period"))
	d.Set("description", instance.DBInstanceDescription)
	d.Set("zone_id", instance.ZoneID)
	d.Set("network_type", instance.NetworkType)

	return nil
}

func resourceAlicloudMongoDBInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient)

	instance, err := client.DescribeMongoDBInstanceById(d.Id(), getRegionId(d, meta))
	if err != nil {
		if NotFoundDBInstance(err) {
			return nil
		}
		return fmt.Errorf("Error Describe MongoDB InstanceAttribute: %#v", err)
	}
	if PayType(instance.ChargeType) == PrePaid {
		return fmt.Errorf("At present, 'PrePaid' instance cannot be deleted and must wait it to be expired and release it automatically.")
	}

	request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
	request.RegionId = getRegionId(d, meta)
	request.QueryParams["DBInstanceId"] = d.Id()
	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		err := client.DeleteMongoDBInstance(request)
		if err != nil {
			if NotFoundDBInstance(err) {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("Delete MongoDB instance timeout and got an error: %#v.", err))
		}

		instance, err := client.DescribeMongoDBInstanceById(d.Id(), getRegionId(d, meta))
		if err != nil {
			if NotFoundError(err) || IsExceptedError(err, InvalidDBInstanceIdNotFound) {
				return nil
			}
			return resource.NonRetryableError(fmt.Errorf("Error Describe MongoDB InstanceAttribute: %#v", err))
		}
		if instance == nil {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("Delete MongoDB instance timeout and got an error: %#v.", err))
	})
}

func buildMongoDBCreateRequest(d *schema.ResourceData, meta interface{}) (*requests.CommonRequest, error) {
	request := CommonRequestInit(getRegionId(d, meta), MONGODBCode, MongoDBDomain)
	request.RegionId = getRegionId(d, meta)
	request.QueryParams["Engine"] = "MongoDB"
	request.QueryParams["EngineVersion"] = d.Get("engine_version").(string)
	request.QueryParams["StorageEngine"] = d.Get("storage_engine").(string)
	request.QueryParams["DBInstanceClass"] = d.Get("instance_class").(string)
	request.QueryParams["DBInstanceStorage"] = strconv.Itoa(d.Get("instance_storage").(int))
	request.QueryParams["DBInstanceDescription"] = d.Get("description").(string)
	request.QueryParams["DBInstanceDescription"] = d.Get("description").(string)
	request.QueryParams["AccountPassword"] = d.Get("password").(string)
	request.QueryParams["ChargeType"] = d.Get("instance_charge_type").(string)

	request.QueryParams["SecurityIPList"] = LOCAL_HOST_IP
	if len(d.Get("security_ips").(*schema.Set).List()) > 0 {
		request.QueryParams["SecurityIPList"] = strings.Join(expandStringList(d.Get("security_ips").(*schema.Set).List())[:], COMMA_SEPARATED)
	}
	// At present, API supports two charge options about 'Prepaid'.
	// 'Month': valid period ranges [1-9]; 'Year': valid period range [1-3]
	// This resource only supports to input Month period [1-9, 12, 24, 36] and the values need to be converted before using them.
	if PayType(request.QueryParams["ChargeType"]) == PrePaid {
		period := d.Get("period").(int)
		request.QueryParams["UsedTime"] = strconv.Itoa(period)
		request.QueryParams["Period"] = string(Month)
		if period > 9 {
			request.QueryParams["UsedTime"] = strconv.Itoa(period / 12)
			request.QueryParams["Period"] = string(Year)
		}
	}

	if zone, ok := d.GetOk("zone_id"); ok && Trim(zone.(string)) != "" {
		request.QueryParams["ZoneId"] = Trim(zone.(string))
	}

	vswitchId := Trim(d.Get("vswitch_id").(string))

	request.QueryParams["NetworkType"] = string(Classic)

	if vswitchId != "" {
		request.QueryParams["VSwitchId"] = vswitchId
		request.QueryParams["NetworkType"] = strings.ToUpper(string(Vpc))

		// check vswitchId in zone
		vsw, err := meta.(*AliyunClient).DescribeVswitch(vswitchId)
		if err != nil {
			return nil, fmt.Errorf("DescribeVSwitche got an error: %#v.", err)
		}

		if request.QueryParams["ZoneId"] == "" {
			request.QueryParams["ZoneId"] = vsw.ZoneId
		} else if strings.Contains(request.QueryParams["ZoneId2"], MULTI_IZ_SYMBOL) {
			zonestr := strings.Split(strings.SplitAfter(request.QueryParams["ZoneId"], "(")[1], ")")[0]
			if !strings.Contains(zonestr, string([]byte(vsw.ZoneId)[len(vsw.ZoneId)-1])) {
				return nil, fmt.Errorf("The specified vswitch %s isn't in the multi zone %s.", vsw.VSwitchId, request.QueryParams["ZoneId"])
			}
		} else if request.QueryParams["ZoneId"] != vsw.ZoneId {
			return nil, fmt.Errorf("The specified vswitch %s isn't in the zone %s.", vsw.VSwitchId, request.QueryParams["ZoneId"])
		}

		request.QueryParams["VpcId"] = vsw.VpcId
	}

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		uuid = resource.UniqueId()
	}
	request.QueryParams["ClientToken"] = fmt.Sprintf("TF-%d-%s", time.Now().Unix(), uuid)

	return request, nil
}
