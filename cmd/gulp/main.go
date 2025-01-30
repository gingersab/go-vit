package main

import (
	"bufio"
	"context"
	"gulp/internal/mon/core"
	"gulp/internal/mon/models"
	"os"
	"time"
)

func main() {
	var sra models.SystemResourceAcquirer = models.InitResourceAcquirer()
	var rm core.Monitor = core.InitResourceMonitor()
	ctx, cancel := context.WithCancel(context.Background())
	dur := 1000 * time.Millisecond
	rm.Start(ctx, sra, dur)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()
}
