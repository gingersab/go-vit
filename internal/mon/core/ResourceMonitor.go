package core

import (
	"context"
	"go-vit/internal/pkg/logfmt"
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
		logfmt.Info.Println("Initiating resource monitoring")
		for {
			select {
			case <-interval.C:
				cpu, err := sre.AcquireCPU()
				if err != nil {
					logfmt.Error.Println(err)
				}
				mem, err := sre.AcquireMem()
				if err != nil {
					logfmt.Error.Println(err)
				}
				driveInfo, err := sre.AcquireDisc()
				if err != nil {
					logfmt.Error.Println(err)
				}
				logfmt.Info.Printf("CPU usage: %.2f%%\n", cpu)
				logfmt.Info.Printf("Memory usage: %.2f%%\n", mem)
				logfmt.Info.Printf("Current drive: %s\n", driveInfo.CDrive)
				logfmt.Info.Printf("Drive mount: %s\n", driveInfo.Mount)
				logfmt.Info.Printf("Filesystem: %s\n", driveInfo.Fs)
				logfmt.Info.Printf("Total space: %d\n", driveInfo.Total)
				logfmt.Info.Printf("Free space: %d\n", driveInfo.Free)
				logfmt.Info.Printf("Used: %d\n", driveInfo.Used)
				logfmt.Info.Printf("Used percent: %.2f%%\n", driveInfo.Perc)

			case <-ctx.Done():
				interval.Stop()
				logfmt.Info.Println("Stopping resource monitoring")
				return
			}
		}
	}()
}
