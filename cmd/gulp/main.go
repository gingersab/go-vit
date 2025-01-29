package main

import (
	"gulp/internal/mon/interfaces"
	"gulp/internal/mon/models"
	"log"
)

func main() {

	var sra interfaces.SystemResourceAcquirer = models.InitResourceAcquirer()
	cpu := sra.AcquireCPU()
	mem := sra.AcquireMem()
	log.Printf("CPU usage: %.2f%%\n", cpu)
	log.Printf("Memory usage: %.2f%%\n", mem)
}
