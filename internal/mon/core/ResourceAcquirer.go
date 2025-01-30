package core

import (
	"fmt"
	"gulp/internal/mon/models"
	"gulp/internal/pkg/gulplog"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemResourceAcquirer interface {
	AcquireCPU() float64
	AcquireMem() float64
	AcquireDisc() (*models.DriveInfo, error)
}

type ResourceAcquirer struct{}

func InitResourceAcquirer() *ResourceAcquirer {
	return &ResourceAcquirer{}
}

func (ResourceAcquirer) AcquireCPU() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		gulplog.Info.Printf("Failed to retrieve CPU usage statistics: %s", err)
	}
	return float64(int(percent[0]*100)) / 100
}

func (ResourceAcquirer) AcquireDisc() (*models.DriveInfo, error) {
	// Step 1: Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		gulplog.Error.Println(err)
	}

	// Step 2: Get the list of partitions
	partitions, err := disk.Partitions(true) // true for detailed info
	if err != nil {
		gulplog.Error.Println(err)
	}

	// Step 3: Find the partition that corresponds to the current working directory
	for _, part := range partitions {
		if currentDir == part.Mountpoint || currentDir[:len(part.Mountpoint)] == part.Mountpoint {
			// Match found, print the drive name (device) and its usage stats
			usage, err := disk.Usage(part.Mountpoint)
			if err != nil {
				gulplog.Error.Println(err)
			}

			d := models.DriveInfo{
				CDrive: part.Device,
				Mount:  part.Mountpoint,
				Fs:     part.Fstype,
				Total:  usage.Total,
				Used:   usage.Used,
				Free:   usage.Free,
				Perc:   usage.UsedPercent,
			}
			return &d, nil
		}
	}
	return nil, fmt.Errorf("failed to identify currently active drive")
}

func (ResourceAcquirer) AcquireMem() float64 {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		gulplog.Info.Printf("Failed to retrieve virtual memory usage statistics: %s", err)
	}
	return float64(int(vmem.UsedPercent*100)) / 100
}
