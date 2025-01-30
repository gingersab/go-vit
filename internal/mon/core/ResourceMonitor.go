package core

import (
	"gulp/internal/mon/models"
	"gulp/internal/pkg/gulplog"
	"time"
)

type Monitor interface {
	Start(models.SystemResourceAcquirer, time.Duration)
	Stop()
}

type ResourceMonitor struct {
	stopChan chan struct{}
}

func InitResourceMonitor() *ResourceMonitor {
	return &ResourceMonitor{
		stopChan: make(chan struct{}),
	}
}

func (rm *ResourceMonitor) Start(sre models.SystemResourceAcquirer, freq time.Duration) {
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
			case <-rm.stopChan:
				return
			}
		}
	}()
}

func (rm *ResourceMonitor) Stop() {
	gulplog.Info.Println("Stopping resource monitoring")
	close(rm.stopChan)
}
