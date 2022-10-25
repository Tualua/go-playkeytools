package goplaykeytools

import (
	"encoding/xml"
	"fmt"
)

type MemUnit string

const (
	MemGiB MemUnit = "GiB"
	MemMiB MemUnit = "MiB"
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
	Cpu    int `xml:"Cpu"`
	Memory PkVmMemSize
}

type PkVmMinConfig struct {
	XMLName xml.Name `xml:"Minimal"`
	PkGsCpuMemConfig
}

type PkVmAutoConfig struct {
	XMLName xml.Name `xml:"VmAutoconf"`
	Minimal PkVmMinConfig
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
	Unit    MemUnit  `xml:"unit,attr"`
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
	RemotePort      int        `xml:"RemotePort"`
	AdapterName     string     `xml:"AdapterName"`
	TemplateFile    string     `xml:"TemplateFile"`
	FilebeatConfig  string     `xml:"FilebeatConfig"`
	LogstashAddress string     `xml:"LogstashAddress"`
	CopyFolder      string     `xml:"CopyFolder"`
	VmAutoconf      PkVmAutoConfig
	PkGsHostCpuMem
}

type PkGsVmsConfig struct {
	XMLName xml.Name        `xml:"Servers"`
	Vms     []PkVmAdvConfig `xml:"Server"`
}

func NewPkGsConfig(hostConfig PkGsHostConfig, vms PkGsVmsConfig) (conf PkGsConfig) {
	conf.HostConfig = hostConfig
	conf.Vms = vms
	return
}

func NewPkGsHostConfig(hostName string, zfsApi PkGsZfsApi, targetAddress string, targetPort int, templateFile string, vmAutoConf PkVmAutoConfig, hostCpus int, hostMemGiB int) (hostConfig PkGsHostConfig) {
	hostConfig.Name = hostName
	hostConfig.ZfsApi = zfsApi
	hostConfig.TargetAddress = targetAddress
	hostConfig.TargetPort = targetPort
	hostConfig.TemplateFile = templateFile
	hostConfig.VmAutoconf = vmAutoConf
	hostConfig.PlaykeyApi = "http://api.playkey.net/"
	hostConfig.RemoteHost = "20.61.216.22"
	hostConfig.RemotePort = 13001
	hostConfig.LogstashAddress = "logstash.playkey.net:12122"
	hostConfig.AdapterName = "NVIDIA GeForce GTX 1080 Ti"
	hostConfig.FilebeatConfig = "/usr/local/share/GameServer/logstash/filebeat.yml"
	hostConfig.CopyFolder = "/home/gamer/vms"
	hostConfig.Cpu = hostCpus
	hostConfig.Memory.Size = hostMemGiB
	hostConfig.Memory.Unit = MemGiB
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

func NewPkGsVmAutoConfig(minCpu int, minMem int, maxCpu int, maxMem int, memUnit MemUnit) (vmAutoConf PkVmAutoConfig) {
	vmAutoConf.Minimal.Cpu = minCpu
	vmAutoConf.Minimal.Memory = PkVmMemSize{Size: minMem, Unit: memUnit}
	vmAutoConf.Cpu = maxCpu
	vmAutoConf.Memory = PkVmMemSize{Size: maxMem, Unit: memUnit}
	return
}

func NewPkVm(name string, numCpu int, mem int, memUnit MemUnit, gpuAddr string, ip string) (vm PkVmAdvConfig) {
	vm.Name = name
	vm.Cpu = numCpu
	vm.Memory = PkVmMemSize{Size: 16, Unit: MemGiB}
	vm.Gpu = gpuAddr
	vm.IP = ip
	return
}
