data "alicloud_mongodb_instances" "mongo" {
  output_file = "out.dat"
}

resource "alicloud_mongodb_instance" "mymongo" {
  instance_class   = "dds.mongo.mid"
  instance_storage = "20"
  engine_version   = "3.4"
  description      = "foobar"
  security_ips     = ["127.0.0.1", "2.2.2.2"]
}

resource "alicloud_mongodb_backup_policy" "mongodb_backup" {
  instance_id             = "${alicloud_mongodb_instance.mongodb.id}"
  preferred_backup_time   = "00:00Z-04:00Z"
  preferred_backup_period = "Friday"
}
