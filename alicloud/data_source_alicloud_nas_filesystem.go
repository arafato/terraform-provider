package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/nas"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAlicloudNASFilesystems() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudNASFilesystemsRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed values
			"filesystems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metered_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mount_targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vswitch_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"access_group": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"network_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mount_target_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"packages": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pakage_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceAlicloudNASFilesystemsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AliyunClient).nasconn

	args := nas.CreateDescribeFileSystemsRequest()
	args.RegionId = getRegionId(d, meta)
	args.PageSize = requests.NewInteger(PageSizeLarge)

	var dbi []nas.FileSystem

	fileSystemID := d.Get("id")

	for {
		resp, err := conn.DescribeFileSystems(args)
		if err != nil {
			return err
		}

		if resp == nil || len(resp.FileSystems.FileSystem) < 1 {
			break
		}

		if fileSystemID != "" {
			for i := range resp.FileSystems.FileSystem {
				if fileSystemID == resp.FileSystems.FileSystem[i].FileSystemId {
					dbi = append(dbi, resp.FileSystems.FileSystem[i])
				}
			}
		} else {
			dbi = append(dbi, resp.FileSystems.FileSystem...)
		}

		if len(resp.FileSystems.FileSystem) < PageSizeLarge {
			break
		}

		if page, err := getNextpageNumber(args.PageNumber); err != nil {
			return err
		} else {
			args.PageNumber = page
		}
	}

	return nasFileSystemDescription(d, dbi)
}

func nasFileSystemDescription(d *schema.ResourceData, dbi []nas.FileSystem) error {
	var ids []string
	var s []map[string]interface{}

	for _, item := range dbi {
		mapping := map[string]interface{}{
			"id":            item.FileSystemId,
			"region_id":     item.RegionId,
			"create_time":   item.CreateTime,
			"protocol_type": item.ProtocolType,
			"storage_type":  item.StorageType,
			"metered_size":  item.MeteredSize,
			"description":   item.Destription,
			"mount_targets": item.MountTargets.MountTarget,
			"packages":      item.Packages.Package,
		}

		ids = append(ids, item.FileSystemId)
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("filesystems", s); err != nil {
		return err
	}

	// create a json file in current directory and write data source to it
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}
	return nil
}
