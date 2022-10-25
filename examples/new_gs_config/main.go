package main

import (
	"encoding/xml"
	"fmt"
	"os"

	pk "github.com/Tualua/go-playkeytools"
)

func main() {
	systemDisk := pk.NewPkGsDisk("data2/kvm/desktop/desktop-master-win10", "iqn.2016-04.net.playkey.iscsi:desktop-", "data2/kvm/desktop", true)
	gamesDisk := pk.NewPkGsDisk("data/reference", "iqn.2016-04.net.playkey.iscsi:games-", "", false)
	storeDisk := pk.NewPkGsDisk("data/reference-store", "iqn.2016-04.net.playkey.iscsi:store-", "", false)
	zfsApi := pk.NewPkGsZfsApi("192.168.255.2", systemDisk, gamesDisk, storeDisk)
	vmAutoConfig := pk.NewPkGsVmAutoConfig(4, 8, 5, 12, pk.MemGiB)
	hostConfig := pk.NewPkGsHostConfig("hv1.EXAMPLE", zfsApi, "192.168.255.2", 3260, "/usr/local/etc/gameserver/template.xml", vmAutoConfig, 6, 12)
	var vms pk.PkGsVmsConfig
	gpus := []string{"0000:0d:00.0", "", "0000:0e:00.0", "0000:03:00.0"}

	for i := 1; i <= 3; i++ {
		vm := pk.NewPkVm(
			fmt.Sprintf("vm%d", i),
			8,
			16,
			pk.MemGiB,
			gpus[i-1],
			fmt.Sprintf("192.168.100.%d", i),
		)
		vms.Vms = append(vms.Vms, vm)
	}
	gsConfig := pk.NewPkGsConfig(hostConfig, vms)

	xmlConfig, _ := xml.MarshalIndent(gsConfig, "", "    ")
	os.WriteFile("conf.xml", xmlConfig, 0644)
}
