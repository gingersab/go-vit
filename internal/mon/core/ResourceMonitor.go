package core

import (
	"context"
	"gulp/internal/pkg/gulplog"
	"time"
)

type Monitor interface {
	Start(context.Context, SystemResourceAcquirer, time.Duration)
}

type ResourceMonitor struct {
}

func InitResourceMonitor() *ResourceMonitor {
	return &ResourceMonitor{}
}

func (rm *ResourceMonitor) Start(ctx context.Context, sre SystemResourceAcquirer, freq time.Duration) {
	interval := time.NewTicker(freq)

	go func() {
		gulplog.Info.Println("Initiating resource monitoring")
		for {
			select {
			case <-interval.C:
				cpu, err := sre.AcquireCPU()
				if err != nil {
					gulplog.Error.Println(err)
				}
				mem, err := sre.AcquireMem()
				if err != nil {
					gulplog.Error.Println(err)
				}
				driveInfo, err := sre.AcquireDisc()
				if err != nil {
					gulplog.Error.Println(err)
				}
				gulplog.Info.Printf("CPU usage: %.2f%%\n", cpu)
				gulplog.Info.Printf("Memory usage: %.2f%%\n", mem)
				gulplog.Info.Printf("Current drive: %s\n", driveInfo.CDrive)
				gulplog.Info.Printf("Drive mount: %s\n", driveInfo.Mount)
				gulplog.Info.Printf("Filesystem: %s\n", driveInfo.Fs)
				gulplog.Info.Printf("Total space: %d\n", driveInfo.Total)
				gulplog.Info.Printf("Free space: %d\n", driveInfo.Free)
				gulplog.Info.Printf("Used: %d\n", driveInfo.Used)
				gulplog.Info.Printf("Used percent: %.2f%%\n", driveInfo.Perc)

			case <-ctx.Done():
				interval.Stop()
				gulplog.Info.Println("Stopping resource monitoring")
				return
			}
		}
	}()
}
