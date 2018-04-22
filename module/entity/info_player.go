package entity

type PlayerInfo struct {
	Id         string
	Ips        []string
	Hostname   string
	Memory     uint64
	LeftMemory uint64
	Cpu        int
	LeftCPU    int
	Status     int
}
