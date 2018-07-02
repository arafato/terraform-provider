data "alicloud_mongodb_instances" "mongo" {
  output_file = "out.dat"
}

resource "alicloud_mongodb_instance" "mymongo" {
  instance_class   = "dds.mongo.mid"
  instance_storage = "10"
  engine_version   = "3.4"
  description      = "foobar"
}
