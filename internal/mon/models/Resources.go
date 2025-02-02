package models

type DriveInfo struct {
	CDrive string
	Mount  string
	Fs     string
	Total  uint64
	Used   uint64
	Free   uint64
	Perc   float64
}

type ResourceStats struct {
	Drive DriveInfo
	Cpu   float64
	Mem   float64
}
