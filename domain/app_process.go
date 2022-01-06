package domain

type MProcess struct {
	Name       string
	Pid        int
	Type       string
	SessionNum int
	Memory     int
	MemoryUnit string
}
