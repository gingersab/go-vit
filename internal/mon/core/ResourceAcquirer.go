package core

import (
	"fmt"
	"go-vit/internal/mon/models"
	"go-vit/internal/pkg/logfmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemResourceAcquirer interface {
	AcquireCPU() (float64, error)
	AcquireMem() (float64, error)
	AcquireDisc() (*models.DriveInfo, error)
}

type ResourceAcquirer struct{}

func InitResourceAcquirer() *ResourceAcquirer {
	return &ResourceAcquirer{}
}

func (ResourceAcquirer) AcquireCPU() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		logfmt.Error.Print(err)
		return 0.0, fmt.Errorf("failed to retrieve CPU statistics")
	}
	return float64(int(percent[0]*100)) / 100, nil
}

func (ResourceAcquirer) AcquireDisc() (*models.DriveInfo, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		logfmt.Error.Println(err)
	}

	partitions, err := disk.Partitions(true)
	if err != nil {
		logfmt.Error.Println(err)
	}

	// Find the partition that corresponds to the current working directory
	for _, part := range partitions {
		if currentDir == part.Mountpoint || currentDir[:len(part.Mountpoint)] == part.Mountpoint {
			// Retrieve drive information and statistics
			usage, err := disk.Usage(part.Mountpoint)
			if err != nil {
				logfmt.Error.Println(err)
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

func (ResourceAcquirer) AcquireMem() (float64, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		logfmt.Error.Print(err)
		return 0.0, fmt.Errorf("failed to retrieve virtual memory statistics")
	}
	return float64(int(vmem.UsedPercent*100)) / 100, nil
}
