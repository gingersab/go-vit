package core

import (
	"context"
	"go-vit/internal/mon/models"
	"go-vit/internal/pkg/logfmt"
	"time"
)

type Monitor interface {
	Start(context.Context, SystemResourceAcquirer, time.Duration) *models.ResourceStats
}

type ResourceMonitor struct {
}

func InitResourceMonitor() *ResourceMonitor {
	return &ResourceMonitor{}
}

func (rm *ResourceMonitor) Start(ctx context.Context, sre SystemResourceAcquirer, freq time.Duration) *models.ResourceStats {
	interval := time.NewTicker(freq)
	stats := &models.ResourceStats{}

	go func() {
		for {
			select {
			case <-interval.C:
				cpu := getCpuStats(sre)
				mem := getMemStats(sre)
				driveInfo := getDriveStats(sre)
				stats.Cpu = cpu
				stats.Drive = *driveInfo
				stats.Mem = mem
			case <-ctx.Done():
				interval.Stop()
				return
			}
		}
	}()
	return stats
}

func getCpuStats(sre SystemResourceAcquirer) float64 {
	cpu, err := sre.AcquireCPU()
	if err != nil {
		logfmt.Error.Println(err)
	}
	return cpu
}

func getMemStats(sre SystemResourceAcquirer) float64 {
	mem, err := sre.AcquireMem()
	if err != nil {
		logfmt.Error.Println(err)
	}
	return mem
}

func getDriveStats(sre SystemResourceAcquirer) *models.DriveInfo {
	driveInfo, err := sre.AcquireDisc()
	if err != nil {
		logfmt.Error.Println(err)
	}
	return driveInfo
}
