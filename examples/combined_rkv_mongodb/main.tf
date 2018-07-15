data "alicloud_rkv_instances" "rkv_instance" {
  output_file = "out.dat"
}

data "alicloud_zones" "default" {
  available_resource_creation = "Rkv"
}

// VPC Resource for Module
resource "alicloud_vpc" "vpc" {
  count      = "${var.vpc_id == "" ? 1 : 0}"
  name       = "${var.vpc_name}"
  cidr_block = "${var.vpc_cidr}"
}

// VSwitch Resource for Module
resource "alicloud_vswitch" "vswitch" {
  count             = "${var.vswitch_id == "" ? 1 : 0}"
  availability_zone = "${var.availability_zone == "" ? data.alicloud_zones.default.zones.0.id : var.availability_zone}"
  name              = "${var.vswitch_name}"
  cidr_block        = "${var.vswitch_cidr}"
  vpc_id            = "${var.vpc_id == "" ? alicloud_vpc.vpc.id : var.vpc_id}"
}

resource "alicloud_rkv_instance" "myredis" {
  instance_class = "${var.instance_class}"
  instance_name  = "${var.instance_name}"
  password       = "${var.password}"
  vswitch_id     = "${var.vswitch_id == "" ? alicloud_vswitch.vswitch.id : var.vswitch_id}"
}

resource "alicloud_rkv_security_ips" "rediswhitelist" {
  instance_id         = "${alicloud_rkv_instance.myredis.id}"
  security_ips        = ["1.1.1.1", "2.2.2.2", "3.3.3.3"]
  security_group_name = "mysecgroup"
}

resource "alicloud_rkv_backup_policy" "redisbackup" {
  instance_id             = "${alicloud_rkv_instance.myredis.id}"
  preferred_backup_time   = "03:00Z-04:00Z"
  preferred_backup_period = ["Monday", "Wednesday", "Friday"]
}

resource "alicloud_mongodb_instance" "mymongo" {
  instance_class   = "dds.mongo.mid"
  instance_storage = "10"
  engine_version   = "3.4"
  description      = "my-description"
  security_ips     = ["127.0.0.1", "2.2.2.2"]
  vswitch_id       = "${var.vswitch_id == "" ? alicloud_vswitch.vswitch.id : var.vswitch_id}"
}

resource "alicloud_mongodb_backup_policy" "mongodb_backup" {
  instance_id             = "${alicloud_mongodb_instance.mymongo.id}"
  preferred_backup_time   = "03:00Z-04:00Z"
  preferred_backup_period = ["Monday", "Wednesday", "Friday"]
}
