package goplaykeytools

import (
	"encoding/xml"
	"fmt"
)

type MemUnit string

const (
	MemGiB = "GiB"
)

type PkGsConfig struct {
	XMLName    xml.Name `xml:"Config"`
	HostConfig PkGsHostConfig
	Vms        PkGsVmsConfig
}

type PkGsZfsApi struct {
	XMLName xml.Name `xml:"ZfsApi"`
	Address string   `xml:"address,attr"`
	Disks   []PkGsDisk
}

type PkGsDisk struct {
	XMLName xml.Name `xml:"Disk"`
	System  int      `xml:"system,attr,omitempty"`
	Origin  string   `xml:"Origin"`
	Prefix  string   `xml:"Prefix"`
	Clone   string   `xml:"Clone,omitempty"`
}

type PkGsCpuMemConfig struct {
	Cpu    string `xml:"Cpu"`
	Memory PkVmMemSize
}

type PkVmMinConfig struct {
	XMLName xml.Name `xml:"Minimal"`
	PkGsCpuMemConfig
}

type PkVmAutoConfig struct {
	XMLName xml.Name `xml:"VmAutoconf"`
	PkGsCpuMemConfig
}

type PkGsHostCpuMem struct {
	PkGsCpuMemConfig
}

type PkVmAdvConfig struct {
	XMLName xml.Name `xml:"Server"`
	Name    string   `xml:"name,attr"`
	PkGsCpuMemConfig
	Gpu string `xml:"Gpu"`
	IP  string `xml:"IP"`
}

type PkVmMemSize struct {
	XMLName xml.Name `xml:"Memory"`
	Unit    string   `xml:"unit,attr"`
	Size    int      `xml:",chardata"`
}

type PkGsHostConfig struct {
	XMLName         xml.Name   `xml:"Host"`
	Name            string     `xml:"name,attr"`
	ZfsApi          PkGsZfsApi `xml:"ZfsApi"`
	PlaykeyApi      string     `xml:"PlaykeyApi"`
	TargetAddress   string     `xml:"TargetAddress"`
	TargetPort      int        `xml:"TargetPort"`
	RemoteHost      string     `xml:"RemoteHost"`
	RemotePort      string     `xml:"RemotePort"`
	AdapterName     string     `xml:"AdapterName"`
	TemplateFile    string     `xml:"TemplateFile"`
	FilebeatConfig  string     `xml:"FilebeatConfig"`
	LogstashAddress string     `xml:"LogstashAddress"`
	CopyFolder      string     `xml:"CopyFolder"`
	VmAutoconf      PkVmAutoConfig
	PkGsHostCpuMem
}

type PkGsVmsConfig struct {
	XMLName xml.Name      `xml:"Servers"`
	Vms     PkVmAdvConfig `xml:"Server"`
}

func NewPkGsConfig(hostConfig PkGsHostConfig, vms PkGsVmsConfig) (conf PkGsConfig) {
	conf.HostConfig = hostConfig
	conf.Vms = vms
	return
}

func NewPkGsHostConfig(zfsApi PkGsZfsApi, targetAddress string, targetPort int, templateFile string, vmAutoConf PkVmAutoConfig) (hostConfig PkGsHostConfig) {
	hostConfig.ZfsApi = zfsApi
	hostConfig.TargetAddress = targetAddress
	hostConfig.TargetPort = targetPort
	hostConfig.TemplateFile = templateFile
	hostConfig.VmAutoconf = vmAutoConf
	return
}

func NewPkGsZfsApi(zfsApiAddr string, systemDisk PkGsDisk, gamesDisk PkGsDisk, storeDisk PkGsDisk) (zfsApi PkGsZfsApi) {
	zfsApi.Address = fmt.Sprintf("http://%s/api", zfsApiAddr)
	zfsApi.Disks = append(zfsApi.Disks, systemDisk)
	zfsApi.Disks = append(zfsApi.Disks, gamesDisk)
	zfsApi.Disks = append(zfsApi.Disks, storeDisk)
	return
}

func NewPkGsDisk(origin string, prefix string, clone string, systemDisk bool) (disk PkGsDisk) {
	if systemDisk {
		disk.System = 1
	}
	if len(clone) > 0 {
		disk.Clone = clone
	}
	disk.Origin = origin
	disk.Prefix = prefix
	return
}

func NewPkVm(name string, numCpu int, mem int, memUnit MemUnit, gpuAddr string, ip string) (vm PkVmAdvConfig) {
	vm.Name = name
	vm.Cpu = fmt.Sprint(numCpu)
	vm.Memory = PkVmMemSize{Size: 16, Unit: MemGiB}
	vm.Gpu = gpuAddr
	vm.IP = ip
	return
}