package system

import (
	pscpu "github.com/shirou/gopsutil/cpu"
	"log"
)

type HostInfo struct {
	Cpu CpuInfo
	Os  OsInfo
}

type OsInfo struct {
	Hostname string
	Platform string
	Version  string
	Kernel   string
}

type CpuInfo struct {
	Architecture   string
	AvailableCores []CpuCore
	EnabledCores   []CpuCore
}

type CpuCore struct {
	Index        int32
	VendorID     string
	ModelName    string
	MaxFrequency float64
}

func (c *CpuInfo) GetCPUInfo() {

	activeCPUInfo, err := pscpu.Info()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(pscpu)

	//c.Architecture = activeCPUInfo.Architecture

	for _, v := range activeCPUInfo {
		core := CpuCore{Index: v.CPU, VendorID: v.VendorID, MaxFrequency: v.Mhz, ModelName: v.ModelName}
		c.EnabledCores = append(c.EnabledCores, core)
	}

}

func GetInfo() HostInfo {
	cpuInfo := CpuInfo{}
	cpuInfo.GetCPUInfo()
	sysInfo := HostInfo{Cpu: cpuInfo}
	return sysInfo
}

func GetSystemDirectories() ([]map[string]string, error) {
	return []map[string]string{}, nil
}

func GetProcessId() (int, error) {
	return 0, nil
}



func EscalatePrivileges() (bool, error) {
	return false, nil
}