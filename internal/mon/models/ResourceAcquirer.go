package models

import (
	"gulp/internal/pkg/log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type ResourceAcquirer struct{}

func InitResourceAcquirer() *ResourceAcquirer {
	return &ResourceAcquirer{}
}

func (ResourceAcquirer) AcquireCPU() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Error.Println("Failed to retrieve CPU usage statistics: ", err)
	}
	return float64(int(percent[0]*100)) / 100
}

func (ResourceAcquirer) AcquireNetwork() {

}

func (ResourceAcquirer) AcquireDisc() {}

func (ResourceAcquirer) AcquireMem() float64 {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		log.Error.Println("Failed to retrieve virtual memory usage statistics: ", err)
	}
	return float64(int(vmem.UsedPercent*100)) / 100
}
