package main

import (
	"bufio"
	"context"
	"go-vit/internal/mon/core"
	"os"
	"time"
)

func main() {
	var sra core.SystemResourceAcquirer = core.InitResourceAcquirer()
	var rm core.Monitor = core.InitResourceMonitor()
	ctx, cancel := context.WithCancel(context.Background())
	dur := 1000 * time.Millisecond
	rm.Start(ctx, sra, dur)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()
	time.Sleep(dur)
}
