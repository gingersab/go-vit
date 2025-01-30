package main

import (
	"bufio"
	"gulp/internal/mon/core"
	"gulp/internal/mon/models"
	"os"
	"time"
)

func main() {
	sra := models.InitResourceAcquirer()
	rm := core.InitResourceMonitor()
	dur := 1000 * time.Millisecond
	rm.Start(sra, dur)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	rm.Stop()
}
