package structComp

type comp struct {
	disk_volume uint32
	disk_name   string
	disk_type   string
}

func (c comp) Getvolume() uint32 {
	return c.disk_volume
}

func (c *comp) Changename(disk_name string) string {
	c.disk_name = disk_name
	return c.disk_name
}
func (c *comp) Changetype(disk_type string) string {
	c.disk_type = disk_type
	return c.disk_type
}
func Computer(disk_volume uint32, disk_name string, disk_type string) comp {
	return comp{
		disk_volume: disk_volume,
		disk_name:   disk_name,
		disk_type:   disk_type,
	}
}
