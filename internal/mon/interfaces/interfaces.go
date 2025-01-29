package interfaces

type SystemResourceAcquirer interface {
	AcquireCPU() float64
	AcquireMem() float64
	AcquireDisc()
	AcquireNetwork()
}
