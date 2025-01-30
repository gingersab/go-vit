package core

import (
	"context"
	"gulp/internal/mon/models"
	"gulp/internal/pkg/gulplog"
	"time"
)

type Monitor interface {
	Start(context.Context, models.SystemResourceAcquirer, time.Duration)
}

type ResourceMonitor struct {
}

func InitResourceMonitor() *ResourceMonitor {
	return &ResourceMonitor{}
}

func (rm *ResourceMonitor) Start(ctx context.Context, sre models.SystemResourceAcquirer, freq time.Duration) {
	interval := time.NewTicker(freq)

	go func() {
		gulplog.Info.Println("Initiating resource monitoring")
		for {
			select {
			case <-interval.C:
				cpu := sre.AcquireCPU()
				mem := sre.AcquireMem()
				gulplog.Info.Printf("CPU usage: %.2f%%\n", cpu)
				gulplog.Info.Printf("Memory usage: %.2f%%\n", mem)
			case <-ctx.Done():
				interval.Stop()
				gulplog.Info.Println("Stopping resource monitoring")
				return
			}
		}
	}()
}
